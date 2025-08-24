terraform {
  required_version = "~> 1.0.0"
  backend "local" {}
}

variable "global_environment_name" {
  description = "A globally unique environment name for S3 buckets."
  type        = string
}

variable "frontend_host" {
  description = "The FQDN to use as the frontend host for the app, i.e. http://mydomain.net"
  type        = string
  default     = ""
}

variable "tls_secret_name" {
  description = "The name of the k8s secret to use for SSL/"
  type        = string
  default     = "wandb-ssl-cert"
}

variable "region" {
  description = "The Azure region in which to place the resources."
  type        = string
  default     = "westus2"
}

variable "license" {
  description = "The W&B license string for your local instance."
  type        = string
}

variable "db_password" {
  description = "Password for the database instance. NOTE: Database is not publicly accessible by default."
  default     = "ComplEx!_DBpAzz32$"
  type        = string
}

variable "wandb_version" {
  description = "The version of wandb to deploy."
  type        = string
  default     = "0.22.0"
}

variable "deployment_is_private" {
  description = "If true, the load balancer will be placed in a private subnet."
  type        = bool
  default     = false
}

variable "use_web_application_firewall" {
  description = "If true, we'll provision a web application firewall for increased security.  This should only be used if deployment_is_private is false."
  type        = bool
  default     = false
}

variable "kubernetes_api_is_private" {
  description = "If true, the kubernetes API server endpoint will be private."
  type        = bool
  default     = true
}

variable "ssl_certificate_name" {
  description = "The name of the SSL certificate that's been manually attached to the application gateway."
  type        = string
  default     = null
}

variable "lets_encrypt_email" {
  description = "The email address to use for automated lets encrypt certs."
  type        = string
  default     = "sysadmin@wandb.com"
}

variable "vpc_cidr_block" {
  description = "CIDR block for the VPC."
  type        = string
  default     = "10.10.0.0/16"
}

variable "public_subnet_cidr_blocks" {
  description = "CIDR blocks for the public VPC subnets. Should be a list of 1 CIDR block."
  type        = list(string)
  default     = ["10.10.0.0/24"]
}

variable "private_ip" {
  description = "The IP address to use when the deployment is private, must be in the public_subnet_cidr_blocks"
  type        = string
  default     = "10.10.0.10"
}

variable "private_subnet_cidr_blocks" {
  description = "CIDR blocks for the private VPC subnets. Should be a list of 1 CIDR block."
  type        = list(string)
  default     = ["10.10.1.0/24"]
}

variable "firewall_ip_address_allow" {
  description = "List of IP addresses that can access the instance via the API.  Defaults to anyone."
  type        = list(string)
  default     = []
}

variable "managed_k8s" {
  description = "Should we manage the k8s cluster with terraform?"
  type        = bool
  default     = false
}

module "infra" {
  source = "./infra"

  global_environment_name      = var.global_environment_name
  region                       = var.region
  db_password                  = var.db_password
  deployment_is_private        = var.deployment_is_private
  kubernetes_api_is_private    = var.kubernetes_api_is_private
  private_ip                   = var.private_ip
  use_web_application_firewall = var.use_web_application_firewall
  vpc_cidr_block               = var.vpc_cidr_block
  public_subnet_cidr_blocks    = var.public_subnet_cidr_blocks
  private_subnet_cidr_blocks   = var.private_subnet_cidr_blocks
  firewall_ip_address_allow    = var.firewall_ip_address_allow
}

locals {
  frontend_host         = coalesce(var.frontend_host, "https://${var.global_environment_name}.${var.region}.cloudapp.azure.com")
  k8s_managed           = var.managed_k8s || !var.kubernetes_api_is_private
  priv_k8s_instructions = <<INST
Check the "Private Control Plane" section of README.md and run:
  az aks command invoke -g ${var.global_environment_name} -n ${var.global_environment_name}-k8s -c \"kubectl apply -f wandb.yaml\" -f wandb.yaml
  INST
  extra_instructions    = var.deployment_is_private ? "\nThis deployment is not accessible from the internet.  You must connect to the VPC to access the service.\nSee the \"Private Deployments\" section of README.md for instructions on configuring DNS, SSL, and VPC Peering." : ""
}

module "kube_yaml" {
  source = "./kube_yaml"

  frontend_host         = local.frontend_host
  deployment_is_private = var.deployment_is_private
  ssl_certificate_name  = var.ssl_certificate_name
  tls_secret_name       = var.tls_secret_name
  lets_encrypt_email    = var.lets_encrypt_email
  license               = var.license
  wandb_version         = var.wandb_version
  bucket                = module.infra.blob_container
  bucket_queue          = module.infra.queue
  mysql                 = module.infra.mysql
  storage_account       = module.infra.storage_account
  storage_key           = module.infra.storage_key
}

module "kube" {
  source = "./kube"

  enabled = local.k8s_managed

  license                     = var.license
  wandb_version               = var.wandb_version
  frontend_host               = local.frontend_host
  deployment_is_private       = var.deployment_is_private
  ssl_certificate_name        = var.ssl_certificate_name
  tls_secret_name             = var.tls_secret_name
  kube_cluster_endpoint       = module.infra.kube_cluster_endpoint
  kube_cert_data              = module.infra.kube_cert_data
  kube_client_key             = module.infra.kube_client_key
  kube_cert_ca                = module.infra.kube_cert_ca
  file_storage_container_name = module.infra.blob_container
  file_metadata_queue_name    = module.infra.queue
  database_endpoint           = module.infra.mysql
  azure_storage_account       = module.infra.storage_account
  azure_storage_key           = module.infra.storage_key
}

output "public_ip" {
  value = module.infra.public_ip
}

output "private_ip" {
  value = module.infra.private_ip
}

output "wandb_vpc_id" {
  value = module.infra.virtual_network_id
}

output "host" {
  value = local.frontend_host
}

output "next_steps" {
  value = local.k8s_managed ? "W&B deployed.  Follow the instructions in README.md to setup SSL (run kubectl apply -f cert-issuer.yaml)${local.extra_instructions}" : "${local.priv_k8s_instructions}${local.extra_instructions}"
}
