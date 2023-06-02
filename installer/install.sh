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
# </Config>


# <ImportInline>
. $DIR/installer/common/kubernetes.sh
. $DIR/installer/common/logging.sh
. $DIR/installer/common/packages.sh
. $DIR/installer/common/semver.sh
. $DIR/installer/common/utils.sh
# </ImportInline>

export PACKAGES=$DIR/packages/kubernetes/$KUBERNETES_VERSION

function main() {
    log_step "Running install with the argument(s): $*"

    require_root_user
    path_add "/usr/local/bin"

    if [ "$AIRGAP" = "1" ]; then
        log_step "Running in airgapped enviroment."
    fi

    must_swap_off
    kubernetes_install_packages
    printf "\n"
}

LOGS_DIR="$DIR/logs"
mkdir -p $LOGS_DIR
LOGFILE="$LOGS_DIR/install-$(date +"%Y-%m-%dT%H-%M-%S").log"

main "$@" 2>&1 | tee $LOGFILE