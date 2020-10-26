#!/bin/bash
#
# This script runs all the necessary terraform commands for you.
set -e

if [ ! -e ".terraform" ]; then
  terraform init
fi

terraform apply -target module.infra -auto-approve && \
terraform apply -target module.kube -auto-approve

echo -e "



------------------------------------------------------------------
\033[0;32mSuccess! Your instance should be online at $(terraform output url)\033[0m
"
