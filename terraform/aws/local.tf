terraform {
  required_version = "~> 0.12.25"
  backend "local" {}
}

variable "global_environment_name" {
  description = "A globally unique environment name for S3 buckets."
  type        = string
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


module "infra" {
  source = "./infra"

  global_environment_name = var.global_environment_name
  db_password             = var.db_password
}

module "kube" {
  source = "./kube"

  license                    = var.license
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
