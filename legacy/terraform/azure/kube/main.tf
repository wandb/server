terraform {
  required_providers {
    kubernetes = "~> 2.8.0"
  }
}


provider "kubernetes" {
  host                   = var.kube_cluster_endpoint
  client_certificate     = var.kube_cert_data
  client_key             = var.kube_client_key
  cluster_ca_certificate = var.kube_cert_ca
}

##########################################
# Variables
##########################################

variable "license" {
  description = "The license string for your local instance."
  type        = string
}

variable "wandb_version" {
  description = "The version of wandb to deploy."
  type        = string
}

variable "kube_cluster_endpoint" {
  description = "The endpoint for the kube cluster."
  type        = string
}

variable "kube_cert_data" {
  description = "Certificate authority data for the kube cluster."
  type        = string
}

variable "kube_client_key" {
  description = "The client key for the kube cluster."
  type        = string
}

variable "kube_cert_ca" {
  description = "The certificate authority for the kube cluster"
  type        = string
}

variable "file_storage_container_name" {
  description = "The name of the file storage container."
  type        = string
}

variable "file_metadata_queue_name" {
  description = "The name of the file metadata queue."
  type        = string
}

variable "azure_storage_account" {
  description = "The storage account name for Azure"
  type        = string
}

variable "azure_storage_key" {
  description = "The storage account key for Azure"
  type        = string
}

variable "frontend_host" {
  description = "The FQDN of the instance"
  type        = string
}

variable "deployment_is_private" {
  description = "If this deployment should be deployed privately"
  type        = bool
}

variable "ssl_certificate_name" {
  description = "The name of the SSL certificate that's been manually attached to the application gateway."
  type        = string
}

variable "tls_secret_name" {
  description = "The name of the secret used for TLS"
  type        = string
}

variable "database_endpoint" {
  description = "The endpoint for the database."
  type        = string
}

variable "enabled" {
  description = "Should we manage k8s with terraform?"
  type        = bool
}

locals {
  host = trimprefix(trimprefix(var.frontend_host, "https://"), "http://")
}

##########################################
# Kubernetes
##########################################

resource "kubernetes_deployment" "wandb" {
  count = var.enabled ? 1 : 0
  metadata {
    name = "wandb"
    labels = {
      app = "wandb"
    }
  }

  spec {
    strategy {
      type = "RollingUpdate"
    }

    replicas = 1

    selector {
      match_labels = {
        app = "wandb"
      }
    }

    template {
      metadata {
        labels = {
          app = "wandb"
        }
      }

      spec {
        container {
          name              = "wandb"
          image             = "wandb/local:${var.wandb_version}"
          image_pull_policy = "Always"

          env {
            name  = "LICENSE"
            value = var.license
          }
          env {
            name  = "BUCKET"
            value = var.file_storage_container_name
          }
          env {
            name  = "BUCKET_QUEUE"
            value = var.file_metadata_queue_name
          }
          env {
            name  = "HOST"
            value = var.frontend_host
          }
          env {
            name  = "MYSQL"
            value = var.database_endpoint
          }
          env {
            name  = "AZURE_STORAGE_ACCOUNT"
            value = var.azure_storage_account
          }
          env {
            name  = "AZURE_STORAGE_KEY"
            value = var.azure_storage_key
          }

          port {
            name           = "http"
            container_port = 8080
            protocol       = "TCP"
          }

          liveness_probe {
            http_get {
              path = "/healthz"
              port = "http"
            }
          }
          readiness_probe {
            http_get {
              path = "/ready"
              port = "http"
            }
          }
          startup_probe {
            http_get {
              path = "/ready"
              port = "http"
            }
            failure_threshold = 60
          }

          resources {
            requests = {
              cpu    = "1500m"
              memory = "4G"
            }
            limits = {
              cpu    = "4000m"
              memory = "8G"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "wandb_service" {
  count = var.enabled ? 1 : 0
  metadata {
    name = "wandb"
  }

  spec {
    selector = {
      app = "wandb"
    }
    port {
      protocol    = "TCP"
      port        = 80
      target_port = 8080
    }
  }
}

resource "kubernetes_ingress" "wandb_ingress" {
  count                  = var.enabled ? 1 : 0
  wait_for_load_balancer = true
  metadata {
    name = "wandb"
    annotations = {
      "kubernetes.io/ingress.class"                       = "azure/application-gateway"
      "appgw.ingress.kubernetes.io/appgw-ssl-certificate" = var.ssl_certificate_name
      "appgw.ingress.kubernetes.io/use-private-ip"        = var.deployment_is_private ? "true" : null
      "cert-manager.io/cluster-issuer"                    = var.deployment_is_private ? null : "issuer-letsencrypt-prod"
      "cert-manager.io/acme-challenge-type"               = var.deployment_is_private ? null : "http01"
    }
  }
  spec {
    tls {
      hosts       = var.deployment_is_private ? null : [local.host]
      secret_name = var.deployment_is_private ? null : var.tls_secret_name
    }
    rule {
      http {
        path {
          path = "/"
          backend {
            service_name = "wandb"
            service_port = 80
          }
        }
      }
    }
    rule {
      host = local.host
      http {
        path {
          path = "/"
          backend {
            service_name = "wandb"
            service_port = 80
          }
        }
      }
    }
  }
}
