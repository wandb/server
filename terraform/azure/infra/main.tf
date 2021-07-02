terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">=2.63.0"
    }
  }
}

provider "azurerm" {
  features {}
}

locals {
  backend_address_pool_name      = "${azurerm_virtual_network.wandb.name}-beap"
  frontend_port_name             = "${azurerm_virtual_network.wandb.name}-feport"
  frontend_ip_configuration_name = "${azurerm_virtual_network.wandb.name}-feip"
  http_setting_name              = "${azurerm_virtual_network.wandb.name}-be-htst"
  listener_name                  = "${azurerm_virtual_network.wandb.name}-httplstn"
  request_routing_rule_name      = "${azurerm_virtual_network.wandb.name}-rqrt"
  redirect_configuration_name    = "${azurerm_virtual_network.wandb.name}-rdrcfg"
  # TODO: this might break if Azure changes the name
  app_gateway_uid_name = "ingressapplicationgateway-${var.global_environment_name}-k8s"

  app_gateway_subnet_name = "${var.global_environment_name}-appgw-subnet"
  k8s_gateway_subnet_name = "${var.global_environment_name}-k8s-subnet"

}

resource "azurerm_resource_group" "wandb" {
  name     = var.global_environment_name
  location = var.region
}

resource "azurerm_virtual_network" "wandb" {
  name                = "${var.global_environment_name}-vpc"
  address_space       = [var.vpc_cidr_block]
  location            = azurerm_resource_group.wandb.location
  resource_group_name = azurerm_resource_group.wandb.name
}

resource "azurerm_subnet" "backend" {
  name                 = local.k8s_gateway_subnet_name
  resource_group_name  = azurerm_resource_group.wandb.name
  virtual_network_name = azurerm_virtual_network.wandb.name
  address_prefixes     = var.private_subnet_cidr_blocks

  service_endpoints                              = ["Microsoft.Sql"] # "Microsoft.Storage", "Microsoft.Web"]
  enforce_private_link_endpoint_network_policies = true
  enforce_private_link_service_network_policies  = true

  # azurerm_subnet_network_security_group_association
}

resource "azurerm_subnet" "frontend" {
  name                 = local.app_gateway_subnet_name
  resource_group_name  = azurerm_resource_group.wandb.name
  virtual_network_name = azurerm_virtual_network.wandb.name
  address_prefixes     = var.public_subnet_cidr_blocks
}

resource "azurerm_public_ip" "wandb" {
  name                = "wandb-public-ip"
  sku                 = "Standard"
  location            = azurerm_resource_group.wandb.location
  resource_group_name = azurerm_resource_group.wandb.name
  allocation_method   = "Static"
  domain_name_label   = var.global_environment_name
}

resource "azurerm_web_application_firewall_policy" "wandb" {
  name                = "wandb-wafpolicy"
  resource_group_name = azurerm_resource_group.wandb.name
  location            = azurerm_resource_group.wandb.location
  tags                = {}


  custom_rules {
    name      = "APIAccessRestrictions"
    priority  = 1
    rule_type = "MatchRule"

    match_conditions {
      transforms = []
      match_variables {
        variable_name = "RemoteAddr"
      }

      operator           = "IPMatch"
      negation_condition = true
      match_values       = length(var.firewall_ip_address_allow) == 0 ? ["10.10.0.0/16"] : var.firewall_ip_address_allow
    }

    match_conditions {
      transforms = []
      match_variables {
        variable_name = "RequestHeaders"
        selector      = "Authorization"
      }

      operator           = "Contains"
      negation_condition = false
      match_values       = ["Basic"]
    }

    action = length(var.firewall_ip_address_allow) == 0 ? "Allow" : "Block"
  }

  policy_settings {
    enabled            = true
    mode               = "Prevention"
    request_body_check = false
  }

  managed_rules {
    managed_rule_set {
      version = "3.2"
    }
  }
}

resource "azurerm_application_gateway" "wandb" {
  name                = "wandb-appgateway"
  resource_group_name = azurerm_resource_group.wandb.name
  location            = azurerm_resource_group.wandb.location

  sku {
    name     = "WAF_v2"
    tier     = "WAF_v2"
    capacity = 2
  }

  gateway_ip_configuration {
    name      = "my-gateway-ip-configuration"
    subnet_id = azurerm_subnet.frontend.id
  }

  frontend_port {
    name = local.frontend_port_name
    port = 80
  }

  frontend_port {
    name = "https_port"
    port = 443
  }

  frontend_ip_configuration {
    name                 = local.frontend_ip_configuration_name
    public_ip_address_id = azurerm_public_ip.wandb.id
  }

  backend_address_pool {
    name = local.backend_address_pool_name
  }

  backend_http_settings {
    name                  = local.http_setting_name
    cookie_based_affinity = "Disabled"
    port                  = 80
    protocol              = "http"
    request_timeout       = 60
  }

  http_listener {
    name                           = local.listener_name
    frontend_ip_configuration_name = local.frontend_ip_configuration_name
    frontend_port_name             = local.frontend_port_name
    protocol                       = "http"
  }

  request_routing_rule {
    name                       = local.request_routing_rule_name
    rule_type                  = "Basic"
    http_listener_name         = local.listener_name
    backend_address_pool_name  = local.backend_address_pool_name
    backend_http_settings_name = local.http_setting_name
  }

  firewall_policy_id = azurerm_web_application_firewall_policy.wandb.id

  depends_on = [
    azurerm_virtual_network.wandb,
    azurerm_public_ip.wandb,
  ]

  lifecycle {
    # K8S will be changing all of these settings so we ignore them.
    # We really only needed this resource to assign a known public IP.
    ignore_changes = [
      ssl_certificate,
      request_routing_rule,
      probe,
      frontend_port,
      http_listener,
      backend_http_settings,
      backend_address_pool,
      tags
    ]
  }
}

data "azurerm_application_gateway" "wandb" {
  name                = "wandb-appgateway"
  resource_group_name = azurerm_resource_group.wandb.name
}


data "azurerm_user_assigned_identity" "wandb" {
  resource_group_name = azurerm_kubernetes_cluster.wandb.node_resource_group
  # The ingress_application_gateway creates a user identity with this name
  # TODO: Figure out how to not rely on this convention
  name = local.app_gateway_uid_name
}

resource "azurerm_role_assignment" "ra3" {
  scope                = azurerm_application_gateway.wandb.id
  role_definition_name = "Contributor"
  # TODO: we can likely use: data.azurerm_application_gateway.wandb.identity.identity_ids ?
  principal_id = data.azurerm_user_assigned_identity.wandb.principal_id
  depends_on = [
    data.azurerm_kubernetes_cluster.wandb,
    data.azurerm_application_gateway.wandb,
  ]
}

resource "azurerm_role_assignment" "ra4" {
  scope                = azurerm_resource_group.wandb.id
  role_definition_name = "Reader"
  # TODO: we can likely use: data.azurerm_application_gateway.wandb.identity.identity_ids ?
  principal_id = data.azurerm_user_assigned_identity.wandb.principal_id
  depends_on = [
    data.azurerm_kubernetes_cluster.wandb,
    data.azurerm_application_gateway.wandb,
  ]
}

resource "azurerm_kubernetes_cluster" "wandb" {
  name                = "${var.global_environment_name}-k8s"
  location            = azurerm_resource_group.wandb.location
  resource_group_name = azurerm_resource_group.wandb.name
  dns_prefix          = var.global_environment_name

  default_node_pool {
    name               = "default"
    node_count         = 2
    vm_size            = "Standard_D4s_v3"
    vnet_subnet_id     = azurerm_subnet.backend.id
    type               = "VirtualMachineScaleSets"
    availability_zones = ["1", "2"]
  }

  network_profile {
    network_plugin    = "azure"
    load_balancer_sku = "standard"
    # TODO: output firewall?
    # load_balancer_profile {
    #   outbound_ip_address_ids = ["${azurerm_public_ip.outbound.id}"]
    # }
  }

  # TODO: RBAC?
  identity {
    type = "SystemAssigned"
  }

  addon_profile {
    http_application_routing {
      enabled = false
    }
    ingress_application_gateway {
      enabled    = true
      gateway_id = azurerm_application_gateway.wandb.id
    }
  }

  private_cluster_enabled = var.kubernetes_api_is_private

  depends_on = [
    azurerm_virtual_network.wandb,
    azurerm_application_gateway.wandb,
  ]
}

data "azurerm_kubernetes_cluster" "wandb" {
  depends_on          = [azurerm_kubernetes_cluster.wandb]
  name                = azurerm_kubernetes_cluster.wandb.name
  resource_group_name = azurerm_resource_group.wandb.name
}

resource "azurerm_mysql_server" "wandb" {
  name                = var.global_environment_name
  location            = azurerm_resource_group.wandb.location
  resource_group_name = azurerm_resource_group.wandb.name

  administrator_login          = "wandb"
  administrator_login_password = var.db_password

  sku_name   = "GP_Gen5_4"
  storage_mb = 5120
  version    = "5.7"

  auto_grow_enabled                 = true
  backup_retention_days             = 14
  geo_redundant_backup_enabled      = false
  infrastructure_encryption_enabled = true
  public_network_access_enabled     = false
  ssl_enforcement_enabled           = true
  ssl_minimal_tls_version_enforced  = "TLS1_2"
}

resource "azurerm_mysql_database" "wandb" {
  name                = "wandb"
  resource_group_name = azurerm_resource_group.wandb.name
  server_name         = azurerm_mysql_server.wandb.name
  charset             = "utf8mb4"
  collation           = "utf8mb4_general_ci"
}

# TODO: can we just use a subnet like?: https://github.com/gustavozimm/terraform-aks-app-gateway-ingress/blob/master/main.tf
resource "azurerm_private_endpoint" "wandb" {
  name                = "wandb-mysql-endpoint"
  location            = azurerm_resource_group.wandb.location
  resource_group_name = azurerm_resource_group.wandb.name
  subnet_id           = azurerm_subnet.backend.id

  private_service_connection {
    name                           = "wandb-mysql-connection"
    private_connection_resource_id = azurerm_mysql_server.wandb.id
    subresource_names              = ["mysqlServer"]
    is_manual_connection           = false
  }
}

data "azurerm_private_endpoint_connection" "wandb" {
  depends_on          = [azurerm_private_endpoint.wandb]
  name                = azurerm_private_endpoint.wandb.name
  resource_group_name = azurerm_resource_group.wandb.name
}

resource "azurerm_storage_account" "wandb" {
  name                     = replace("${var.global_environment_name}-storage", "-", "")
  resource_group_name      = azurerm_resource_group.wandb.name
  location                 = azurerm_resource_group.wandb.location
  account_tier             = "Standard"
  account_replication_type = "ZRS"
  min_tls_version          = "TLS1_2"

  blob_properties {
    cors_rule {
      allowed_headers    = ["*"]
      allowed_methods    = ["GET", "HEAD", "PUT"]
      allowed_origins    = ["*"]
      exposed_headers    = ["ETag"]
      max_age_in_seconds = 3600
    }
  }
}

resource "azurerm_storage_container" "wandb" {
  name                  = "wandb-files"
  storage_account_name  = azurerm_storage_account.wandb.name
  container_access_type = "private"
}

resource "azurerm_storage_queue" "wandb" {
  name                 = "wandb-file-metadata"
  storage_account_name = azurerm_storage_account.wandb.name
}

resource "azurerm_eventgrid_system_topic" "wandb" {
  name                   = "wandb-file-metadata-topic"
  location               = azurerm_resource_group.wandb.location
  resource_group_name    = azurerm_resource_group.wandb.name
  source_arm_resource_id = azurerm_storage_account.wandb.id
  topic_type             = "Microsoft.Storage.StorageAccounts"
}

resource "azurerm_eventgrid_system_topic_event_subscription" "wandb" {
  name = "wandb-file-metadata-subscription"
  # scope                = azurerm_resource_group.wandb.id
  system_topic         = azurerm_eventgrid_system_topic.wandb.name
  resource_group_name  = azurerm_resource_group.wandb.name
  included_event_types = ["Microsoft.Storage.BlobCreated"]
  subject_filter {
    subject_begins_with = "/blobServices/default/containers/wandb-files/blobs/"
  }

  storage_queue_endpoint {
    storage_account_id = azurerm_storage_account.wandb.id
    queue_name         = azurerm_storage_queue.wandb.name
  }
}
