# Deprecated: This terraform module has been deprecated please use our [new module](https://github.com/wandb/terraform-aws-wandb) instead.

## Installing W&B Local in AWS

This repo contains a terraform stack you can use to deploy your very own W&B Local instance into your AWS account.

### Prerequisites

To use this install guide, you must have the following utilities installed already:
* [Terraform 0.12.25](https://releases.hashicorp.com/terraform/0.12.25/)
* [AWS CLI](https://aws.amazon.com/cli/)

### Preparing for Install
For best results, we recommend applying this terraform stack in an empty AWS subaccount. We also recommend against making custom modifications to this terraform stack (outside of specifying configurable variables) -- we can only guarantee that the default stack in the main branch of this repo is fully functional.

First, make sure you can access your AWS account with the AWS CLI. You may need to copy credentials into `~/.aws/credentials` and set your `AWS_PROFILE` environment variable.

Next, create a file in this directory (or wherever you plan to run terraform) called `terraform.tfvars`. In this file, you'll define some variables to configure your install. This file should have at least the following entries:

```
global_environment_name = "<YOUR_UNIQUE_ENVIRONMENT_NAME_HERE>"
license = "<YOUR_LICENSE_HERE>"
```

Finally, run the following command to initialize terraform:
```
terraform init
```

### Private Installs
This stack has an option to deploy W&B entirely inside a VPC and expose nothing to the internet. To choose this option, specify `deployment_is_private = true` in your `terraform.tfvars` file. Note that if you choose this option, you will only be able to access your W&B instance from inside your VPC (or via VPN connection).

This stack also contains an option to make the kubernetes API server endpoint private. To choose this option, specify `kubernetes_api_is_private = true` in your `terraform.tfvars` file. If you choose this option, you'll have to run the "Installing W&B in the Cluster" step below with a VPN connection active, since the kubernetes API server endpoint will be private to the VPC. This stack does not provision any VPN resources -- you will have to set up VPN access manually between the two terraform steps described below.

### Creating the Cluster (and other resources)
To create the cluster, run the following command:
```
terraform apply -target module.infra -auto-approve
```

## Installing W&B in the Cluster
To install W&B in the cluster, run the following command:
```
terraform apply -target module.kube -auto-approve
```
### Alternatively: Quick Install
If you'd like to run all the terraform steps in one go, you can use our included shell script:
```
./install_wandb.sh
```

### (IMPORTANT!) Save TF State
After install, terraform will generate a `terraform.tfstate` file. It is *extremely* important that you do not lose this file. Without this state file, you will no longer be able to manage your W&B install with terraform. `terraform.tfstate` must be present in your working directory whenever you run any terraform commands. We recommend backing this file up to a well known location.

### Cluster Administration
After install, this terraform stack will output a `kubeconfig.yaml` file you can use to administer the cluster. Once the cluster is done installing, try running the following command to see pod status:
```
kubectl --kubeconfig=kubeconfig.yaml get pods
```
