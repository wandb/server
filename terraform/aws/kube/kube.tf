provider "kubernetes" {
  load_config_file = "false"

  host = var.kube_cluster_endpoint

  cluster_ca_certificate = base64decode(var.kube_cert_data)

  config_context_auth_info = "aws"
  config_context_cluster   = "kubernetes"

  exec {
    api_version = "client.authentication.k8s.io/v1alpha1"
    command     = "aws-iam-authenticator"
    args        = ["token", "-i", "wandb"]
  }
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

variable "file_storage_bucket_name" {
  description = "The name of the file storage bucket."
  type        = string
}

variable "file_storage_bucket_region" {
  description = "The region in which the file storage bucket resides."
  type        = string
}

variable "file_metadata_queue_name" {
  description = "The name of the file metadata queue."
  type        = string
}

variable "database_endpoint" {
  description = "The endpoint for the database."
  type        = string
}

##########################################
# Kubernetes
##########################################

resource "kubernetes_deployment" "wandb" {
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
            value = "s3:///${var.file_storage_bucket_name}"
          }
          env {
            name  = "BUCKET_QUEUE"
            value = "sqs://${var.file_metadata_queue_name}"
          }
          env {
            name  = "AWS_REGION"
            value = var.file_storage_bucket_region
          }
          env {
            name  = "MYSQL"
            value = "mysql://${var.database_endpoint}"
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

          resources {
            requests {
              cpu    = "1500m"
              memory = "4G"
            }
            limits {
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
  metadata {
    name = "wandb"
  }

  spec {
    type = "NodePort"
    selector = {
      app = "wandb"
    }
    port {
      port      = 8080
      node_port = 32543
    }
  }
}

resource "local_file" "kubeconfig" {
  filename = "kubeconfig.yaml"
  content  = <<KUBECONFIG
    apiVersion: v1
    kind: Config
    clusters:
    - cluster:
        server: ${var.kube_cluster_endpoint}
        certificate-authority-data: ${var.kube_cert_data}
      name: kubernetes
    contexts:
    - context:
        cluster: kubernetes
        user: aws
      name: aws
    current-context: aws
    preferences: {}
    users:
    - name: aws
      user:
        exec:
          apiVersion: client.authentication.k8s.io/v1alpha1
          command: aws-iam-authenticator
          args:
            - "token"
            - "-i"
            - "wandb"
KUBECONFIG
}
