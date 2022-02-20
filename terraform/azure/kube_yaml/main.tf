variable "frontend_host" {
  description = "The DNS entry to use as the frontend host for the app."
  type        = string
}

variable "wandb_version" {
  description = "The version of wandb to deploy."
  type        = string
  default     = "latest"
}

variable "license" {
  description = "The W&B license string for your local instance."
  type        = string
}

variable "bucket" {
  description = "The object store url"
  type        = string
}

variable "bucket_queue" {
  description = "The bucket queue"
  type        = string
}

variable "mysql" {
  description = "The mysql database"
  type        = string
}

variable "storage_account" {
  description = "Azure storage account"
  type        = string
}

variable "storage_key" {
  description = "Azure storage key"
  type        = string
}

variable "ssl_certificate_name" {
  description = "The name of the SSL certificate that's been manually attached to the application gateway."
  type        = string
}

variable "tls_secret_name" {
  description = "The name of the k8s secret to use for SSL/"
  type        = string
  default     = "wandb-ssl-cert"
}

variable "lets_encrypt_email" {
  description = "The email address to use for automated lets encrypt certs."
  type        = string
  default     = "sysadmin@wandb.com"
}

variable "deployment_is_private" {
  description = "If this deployment should be deployed privately"
  type        = bool
}

locals {
  host                = trimprefix(trimprefix(var.frontend_host, "https://"), "http://")
  private_annotations = <<ANN
    ${var.ssl_certificate_name != null ? "appgw.ingress.kubernetes.io/appgw-ssl-certificate: ${var.ssl_certificate_name}" : ""}
    appgw.ingress.kubernetes.io/use-private-ip: "true"
  ANN
  public_annotations  = <<ANN
    cert-manager.io/cluster-issuer: issuer-letsencrypt-prod
    cert-manager.io/acme-challenge-type: http01
  ANN
  tls                 = <<TLS
  tls:
  - hosts:
    - ${local.host}
    secretName: ${var.tls_secret_name}
  TLS
}

resource "local_file" "wandb_kube" {
  filename = "wandb.yaml"
  content  = <<KUBE
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wandb
  labels:
    app: wandb
spec:
  strategy:
    type: RollingUpdate
  replicas: 1
  selector:
    matchLabels:
      app: wandb
  template:
    metadata:
      labels:
        app: wandb
    spec:
      containers:
        - name: wandb
          env:
            - name: LICENSE
              value: ${var.license}
            - name: BUCKET
              value: ${var.bucket}
            - name: HOST
              value: ${var.frontend_host}
            - name: BUCKET_QUEUE
              value: ${var.bucket_queue}
            - name: MYSQL
              value: ${var.mysql}
            - name: AZURE_STORAGE_ACCOUNT
              value: ${var.storage_account}
            - name: AZURE_STORAGE_KEY
              value: ${var.storage_key}
          imagePullPolicy: Always
          image: wandb/local:${var.wandb_version}
          ports:
            - name: http
              containerPort: 8080
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /healthz
              port: http
          readinessProbe:
            httpGet:
              path: /ready
              port: http
          startupProbe:
            httpGet:
              path: /ready
              port: http
            failureThreshold: 60
          resources:
            requests:
              cpu: "1500m"
              memory: 4G
            limits:
              cpu: "4000m"
              memory: 8G
---
apiVersion: v1
kind: Service
metadata:
  name: wandb
spec:
  selector:
    app: wandb
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: wandb
  annotations:
    kubernetes.io/ingress.class: azure/application-gateway
    ${var.deployment_is_private ? local.private_annotations : local.public_annotations}
spec:
  ${var.deployment_is_private ? "" : local.tls}
  rules:
  - http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: wandb
            port:
              number: 80
  - host: ${local.host}
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: wandb
            port:
              number: 80
KUBE
}

resource "local_file" "wandb_kube_ssl" {
  filename = "cert-issuer.yaml"
  content  = <<KUBE
apiVersion: cert-manager.io/v1alpha2
kind: ClusterIssuer
metadata:
  name: issuer-letsencrypt-prod
spec:
  acme:
    email: ${var.lets_encrypt_email}
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: secret-letsencrypt-prod
    solvers:
    - http01:
        ingress:
            class: azure/application-gateway
KUBE
}
