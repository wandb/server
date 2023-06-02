#!/bin/bash

set -e

DIR=.

# <Config>
export AIRGAP=0
export ARCH="amd64"

export KUBERNETES_VERSION="1.27.2"
export CRICTL_VERSION="1.27.0"
export CNI_PLUGINS_VERSION="1.3.0"
export CONTAINERD_VERSION="1.7.1"
export RUNC_VERSION="1.1.7"

export OPENSSL_VERSION="3.1.1"
# </Config>

# <ImportInline>
. $DIR/installer/common/kubernetes.sh
. $DIR/installer/common/logging.sh
. $DIR/installer/common/discover.sh
. $DIR/installer/common/packages.sh
. $DIR/installer/common/semver.sh
. $DIR/installer/common/utils.sh
. $DIR/installer/common/dependencies.sh
# </ImportInline>

export PACKAGES=$DIR/packages/kubernetes/$KUBERNETES_VERSION
export DEPENDENCIES=$DIR/packages/deps

log_step "Generating installer"
pnpm build:installer:airgap

kubernetes_packages_download
dependencies_download

log_step "Creating tar"
tar -czvf bundle-$ARCH-$KUBERNETES_VERSION.tar.gz $DIR/install.sh $DIR/packages

printf "\n"