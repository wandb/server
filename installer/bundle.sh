#!/bin/bash

set -e

DIR=.

# <ImportInline>
. $DIR/installer/configs/1.27.2.sh
. $DIR/installer/configs/base.sh
. $DIR/installer/common/kubernetes.sh
. $DIR/installer/common/addons.sh
. $DIR/installer/common/kubeadm.sh
. $DIR/installer/common/logging.sh
. $DIR/installer/common/discover.sh
. $DIR/installer/common/packages.sh
. $DIR/installer/common/semver.sh
. $DIR/installer/common/utils.sh
. $DIR/installer/common/dependencies.sh
. $DIR/installer/common/images.sh
# </ImportInline>

log_step "Generating installer"
pnpm build:installer:airgap

kubernetes_packages_download
dependencies_download
images_download
addons_download

log_step "Creating tar"
tar -czvf installer-$ARCH-$KUBERNETES_VERSION.tar.gz $DIR/install.sh $DIR/packages

printf "\n"