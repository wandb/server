output "subnet_id" {
  value = azurerm_kubernetes_cluster.wandb.default_node_pool[0].vnet_subnet_id
}

output "blob_container" {
  value = "az://${azurerm_storage_account.wandb.name}/${azurerm_storage_container.wandb.name}"
}

output "queue" {
  value = "az://${azurerm_storage_account.wandb.name}/${azurerm_storage_queue.wandb.name}"
}

output "mysql" {
  value = "mysql://${urlencode(join("@", [azurerm_mysql_server.wandb.administrator_login, azurerm_mysql_server.wandb.name]))}:${urlencode(azurerm_mysql_server.wandb.administrator_login_password)}@${data.azurerm_private_endpoint_connection.wandb.private_service_connection.0.private_ip_address}/${azurerm_mysql_database.wandb.name}"
}

output "storage_key" {
  value = azurerm_storage_account.wandb.primary_access_key
}

output "storage_account" {
  value = azurerm_storage_account.wandb.name
}

output "public_ip" {
  value = azurerm_public_ip.wandb.ip_address
}

output "public_dns" {
  value = azurerm_public_ip.wandb.fqdn
}

output "resource" {
  value = azurerm_public_ip.wandb.resource_group_name
}

#output "private_ip" {
#  value = data.azurerm_private_endpoint_connection.wandb_web.private_service_connection.0.private_ip_address
#}

output "kube_cluster_endpoint" {
  value = azurerm_kubernetes_cluster.wandb.kube_config.0.host
}

output "kube_cert_data" {
  value = base64decode(azurerm_kubernetes_cluster.wandb.kube_config.0.client_certificate)
}

output "kube_client_key" {
  value = base64decode(azurerm_kubernetes_cluster.wandb.kube_config.0.client_key)
}

output "kube_cert_ca" {
  value = base64decode(azurerm_kubernetes_cluster.wandb.kube_config.0.cluster_ca_certificate)
}

output "identity" {
  value = data.azurerm_user_assigned_identity.wandb.id
}
