terraform {
  required_version = "~> 0.12.25"
  backend "local" {}
}

variable "global_environment_name" {
  description = "A globally unique environment name for S3 buckets."
  type        = string
}

variable "aws_region" {
  description = "The AWS region in which to place the resources."
  type        = string
  default     = "us-west-2"
}

variable "license" {
  description = "The license string for your local instance."
  type        = string
}

variable "db_password" {
  description = "Password for the database instance. NOTE: Database is not publicly accessible by default."
  default     = "wandb_root_password"
  type        = string
}

variable "wandb_version" {
  description = "The version of wandb to deploy."
  type        = string
  default     = "0.9.30"
}

variable "deployment_is_private" {
  description = "If true, the load balancer will be placed in a private subnet, and the kubernetes API server endpoint will be private."
  type        = bool
  default     = false
}

variable "kubernetes_api_is_private" {
  description = "If true, the kubernetes API server endpoint will be private."
  type        = bool
  default     = false
}
variable "vpc_cidr_block" {
  description = "CIDR block for the VPC."
  type        = string
  default     = "10.10.0.0/16"
}

variable "public_subnet_cidr_blocks" {
  description = "CIDR blocks for the public VPC subnets. Should be a list of 2 CIDR blocks."
  type        = list(string)
  default     = ["10.10.0.0/24", "10.10.1.0/24"]
}

variable "private_subnet_cidr_blocks" {
  description = "CIDR blocks for the private VPC subnets. Should be a list of 2 CIDR blocks."
  type        = list(string)
  default     = ["10.10.2.0/24", "10.10.3.0/24"]
}


module "infra" {
  source = "./infra"

  global_environment_name    = var.global_environment_name
  aws_region                 = var.aws_region
  db_password                = var.db_password
  deployment_is_private      = var.deployment_is_private
  kubernetes_api_is_private  = var.kubernetes_api_is_private
  vpc_cidr_block             = var.vpc_cidr_block
  public_subnet_cidr_blocks  = var.public_subnet_cidr_blocks
  private_subnet_cidr_blocks = var.private_subnet_cidr_blocks
}

module "kube" {
  source = "./kube"

  license                    = var.license
  wandb_version              = var.wandb_version
  kube_cluster_endpoint      = module.infra.eks_cluster_endpoint
  kube_cert_data             = module.infra.eks_cert_data
  file_storage_bucket_name   = module.infra.s3_bucket_name
  file_storage_bucket_region = module.infra.s3_bucket_region
  file_metadata_queue_name   = module.infra.sqs_queue_name
  database_endpoint          = module.infra.rds_connection_string
}

output "url" {
  value = module.infra.lb_address
}
