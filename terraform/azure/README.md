# Installing W&B Local in Azure

This repo contains a terraform stack you can use to deploy your very own W&B Local instance into your Azure account.

## Prerequisites

To use this install guide, you must have the following utilities installed already:
* [Terraform >= 0.12.25](https://releases.hashicorp.com/terraform/0.12.25/)
* [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)

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

- lets_encrypt_email - _If you use lets encrypt to handle SSL we'll create a cluster issuer using this email address, see SSL below_
- frontend_host - _By default we'll use the azure provided dns at envname.regionname.cloudapp.azure.com_
- kubernetes_api_is_private - _By default the k8s control is on the internet, to make your installation more secure set this to true, see Private Control Plane below_

## Installation

Be sure you created a file named `terraform.tfvars` in this directory then run:

```
terraform init
terraform apply
```

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

## SSL

If you don't have an SSL certificate we recommend using lets-encrypt.  Install cert manager with:

```
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml
kubectl apply -f cert-issuer.yaml
```

### References

https://github.com/Azure/terraform-azurerm-appgw-ingress-k8s-cluster/blob/master/main.tf
https://github.com/gustavozimm/terraform-aks-app-gateway-ingress/blob/master/main.tf

https://blog.baeke.info/2020/10/25/
https://docs.microsoft.com/en-us/azure/application-gateway/ingress-controller-letsencrypt-certificate-application-gateway