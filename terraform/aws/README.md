# Installing W&B Local in AWS

This repo contains a terraform stack you can use to deploy your very own W&B Local instance into your AWS account.

## Prerequisites

To use this install guide, you must have the following utilities installed already:
* [Terraform 0.12.25 or higher](https://learn.hashicorp.com/tutorials/terraform/install-cli)
* [AWS CLI](https://aws.amazon.com/cli/)
* [aws-iam-authenticator](https://docs.aws.amazon.com/eks/latest/userguide/install-aws-iam-authenticator.html)

## Preparing for Install
For best results, we recommend applying this terraform stack in an empty AWS subaccount. We also recommend against making custom modifications to this terraform stack (outside of specifying configurable variables) -- we can only guarantee that the default stack in the main branch of this repo is fully functional.

First, make sure you can access your AWS account with the AWS CLI. You may need to copy credentials into `~/.aws/credentials` and set your `AWS_PROFILE` environment variable.

Next, create a file in this directory (or wherever you plan to run terraform) called `terraform.tfvars`. In this file, you'll define some variables to cofigure your install. This file should have the following entries:

```
global_environment_name = "<YOUR_UNIQUE_ENVIRONMENT_NAME_HERE>"
license = "<YOUR_LICENSE_HERE>"
```

Finally, run the following command to initialize terraform:
```
terraform init
```

## Creating the Cluster (and other resources)
To create the cluster, run the following command:
```
terraform apply -target module.infra -auto-approve
```

## Installing W&B in the Cluster
To install W&B in the cluster, run the following command:
```
terraform apply -target module.kube -auto-approve
```
## Alternatively: Quick Install
If you'd like to run all the terraform steps in one go, you can use our included shell script:
```
./install_wandb.sh
```

## (IMPORTANT!) Save TF State
After install, terraform will generate a `terraform.tfstate` file. It is *extremely* important that you do not lose this file. Without this state file, you will no longer be able to manager your W&B install with terraform. `terraform.tfstate` must be present in your working directory whenever you run any terraform commands. We recommend backing this file up to a well known location.

## Cluster Administration
After install, this terraform stack will output a `kubeconfig.yaml` file you can use to administer the cluster. Once the cluster is done installing, try running the following command to see pod status:
```
kubectl --kubeconfig=kubeconfig.yaml get pods
```
