# Installing W&B Local in Azure

This repo contains a terraform stack you can use to deploy your very own W&B Local instance into your Azure account.

## Prerequisites

To use this install guide, you must have the following utilities installed already:
* [Terraform >= 0.12.25](https://releases.hashicorp.com/terraform/0.12.25/)
* [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)

## Overview

This terraform will deploy an application gateway and AKS cluster running wandb connected to a blob container and MySQL database.

## Preparing for install

You'll need to ensure your az command is logged in to the right account by running `az login`.  By default we deploy an AKS cluster with the control plane on the internet.  If you want to keep the kubernetes control plane off the internet see Private Control Plane below.

You'll want to store your state in a safe place, here are [instructions for using Azure](https://docs.microsoft.com/en-us/azure/developer/terraform/create-k8s-cluster-with-aks-applicationgateway-ingress#configure-azure-storage-to-store-terraform-state).

## Variables

Check the source code of this repo to see all variables, at a minimum you should be aware the following variables and specify them in a file named `terraform.tfvars`:

### Required

- global_environment_name - _This must be unique to your org as we'll create a storage account / container using this name as well as an azure resource group, eg (my-company-name-wandb)_
- license - _The W&B license key provided by your account executive_
- region - _By default we use westus2, if you would like to use a different region override it here_

### Optional

- db_password - _The password used to connect to your database, you should probably override this_
- lets_encrypt_email - _If you use lets encrypt to handle SSL we'll create a cluster issuer using this email address, see SSL below_
- frontend_host - _By default we'll use the azure provided dns at envname.regionname.cloudapp.azure.com_
- kubernetes_api_is_private - _By default the k8s control is on the internet, to make your installation more secure set this to true, see Private Control Plane below_
- deployment_is_private - _By default we'll provision an IP address that's accessible on the internet.  Adding this flag will make the load balancer only listen on the VPC.  We can't provision SSL certificates in this mode, so you'll have to do it manually._
- ssl_certificate_name - _If you're running the application privately, you'll need to configure an SSL certificate manually.  This variable attaches that certificate to the k8s ingress_
- use_web_application_firewall - _When running the service on the internet, additional security can be provided by enabling a web application firewall. This will incur additional charges_
- firewall_ip_address_allow - _By default we allow any IP address to log data and use our API's.  Providing a list of IP ranges will limit all programmatic access to those IP's.  Access to the UI will still be allowed from anywhere on the internet.  This only works when the web application firewall is enabled_

## Installation

Be sure you created a file named `terraform.tfvars` in this directory then run:

```
terraform init
terraform apply
```

Be sure to save all `terraform.*` files in a safe place.

## SSL

If you don't have an SSL certificate we recommend using [lets-encrypt](https://letsencrypt.org).  Install cert manager with:

```
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.7.1/cert-manager.yaml
kubectl apply -f cert-issuer.yaml
```

> NOTE: You can't use cert manager when the deployment is private, see "Private Deployment SSL".

## Upgrades

Modify the `wandb_version` value in `terraform.tfvars` to the latest version of wandb found in this repositories releases and run `terraform apply`.

### Private Control Plane

The simplest way to communicate with a private kubernetes cluster is using a preview feature in Azure called [AKS Run Command](https://docs.microsoft.com/en-us/azure/aks/private-clusters#aks-run-command-preview).  You can enable it by running:

```
az feature register --namespace "Microsoft.ContainerService" --name "RunCommandPreview"
```

Wait 10-60 seconds, then run: `az provider register --namespace Microsoft.ContainerService`

Alternatively you can use the [Azure VPN](https://docs.microsoft.com/en-us/azure/vpn-gateway/openvpn-azure-ad-tenant) app to connect to the VPN after creating the infra with:

```
terraform init
terraform apply -t module.infra
# Connect to VPN
terraform apply -t module.kube
```

If you can't connect to the VPN we output a `wandb.yaml` k8s manifest in the infra step that you can apply with:

```
az aks command invoke -g $GLOBAL_ENV_NAME -n $GLOBAL_ENV_NAME-k8s -c "kubectl apply -f wandb.yaml" -f wandb.yaml
```

## Private Deployments

Setting `deployment_is_private` will make the load balancer only listen on the internal network.  You'll need to setup DNS and SSL manually.  After the terraform has completed you can use the private_ip to add an A record to your DNS service.

> NOTE: Because of [current limitations](https://github.com/Azure/application-gateway-kubernetes-ingress/issues/741) with the way the Azure Application Gateway integrates with Azure Kubernetes Engine we must provision a public IP address.  When `var.deployment_is_private` we block all internet traffic to the public IP.

### DNS

When the deployment is private all we can provide is an IP address.  For SSL to work, you'll need to configure your internal DNS to resolve the IP address (10.10.0.10 by default).

### Private Deployment SSL

When the application is running privately we can not automatically provision SSL certificates.  You'll need to provision an SSL certificate yourself.  You should obtain a certificate from a trusted provider, using self-signed certificates will cause lot's of angst for your end users.

Once you've obtained a trusted certificate see the [following documentation](https://azure.github.io/application-gateway-kubernetes-ingress/annotations/#appgw-ssl-certificate) for configuring it with your application gateway.  You'll need to set the `var.ssl_certificate_name` variable for it to be associated with your deployment.

### VPC Peering

If you're deploying this resource into an isolated VPC, you'll need to peer it with your existing VPC.
You can find the wandb_vpc_id in the output of `terraform apply`.  Here's example terraform for configuring
VPC peering:

```terraform
resource "azurerm_virtual_network_peering" "wandb" {
  name                      = "my_virtual_network"
  resource_group_name       = "my_resource_group"
  virtual_network_name      = "wandb_network"
  remote_virtual_network_id = outputs.wandb_vpc_id
}
```

## References

- https://azure.github.io/application-gateway-kubernetes-ingress
- https://github.com/Azure/terraform-azurerm-appgw-ingress-k8s-cluster/blob/master/main.tf
- https://github.com/gustavozimm/terraform-aks-app-gateway-ingress/blob/master/main.tf
- https://blog.baeke.info/2020/10/25/
- https://docs.microsoft.com/en-us/azure/application-gateway/ingress-controller-letsencrypt-certificate-application-gateway
- https://docs.microsoft.com/en-us/azure/web-application-firewall/ag/create-custom-waf-rules
- https://github.com/Azure/application-gateway-kubernetes-ingress/issues/741