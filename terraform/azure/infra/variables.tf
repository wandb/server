variable "global_environment_name" {
  description = "A globally unique environment name for S3 buckets."
  type        = string
}

variable "region" {
  description = "The Azure region in which to place the resources."
  type        = string
  default     = "westus2"
}

variable "db_password" {
  description = "Password for the database instance. NOTE: Database is not publicly accessible by default."
  type        = string
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
  description = "CIDR blocks for the public VPC subnets. Should be a list of 1 CIDR blocks."
  type        = list(string)
  default     = ["10.10.0.0/24"]
}

variable "private_subnet_cidr_blocks" {
  description = "CIDR blocks for the private VPC subnets. Should be a list of 1 CIDR blocks."
  type        = list(string)
  default     = ["10.10.1.0/24"]
}

variable "firewall_ip_address_allow" {
  description = "List of IP addresses that can access the instance via the API.  Defaults to anyone."
  type        = list(string)
  default     = []
}
