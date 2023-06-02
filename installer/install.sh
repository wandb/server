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
export HOSTNAME="$(hostname | tr '[:upper:]' '[:lower:]')"

function setup() {
    require_root_user
    path_add "/usr/local/bin"

    if [ "$AIRGAP" = "1" ]; then
        log_step "Running in airgapped enviroment."
    fi

    kubernetes_load_modules
    kubernetes_load_sysctl
    must_swap_off
    kubernetes_install_packages
}

function init() {
    API_SERVICE_ADDRESS="$PRIVATE_ADDRESS:6443"

    set +o pipefail

    # cmd_retry 3 kubeadm init \
    #     --ignore-preflight-errors="all" \
    #     | tee /tmp/kubeadm-init
    
    log_step "Waiting for kubernetes api health to report ok"
    if ! spinner_until 120 kubeadm_api_is_healthy; then
        bail "Kubernetes API failed to report healthy"
    fi
    
    export KUBECONFIG=/etc/kubernetes/admin.conf
    
    kubectl cluster-info
    log_success "Cluster initialized"
}

function main() {
    log_step "Running install with the argument(s): $*"
    
    discover
    # setup    
    init
    
    printf "\n"
}

LOGS_DIR="$DIR/logs"
mkdir -p $LOGS_DIR
LOGFILE="$LOGS_DIR/install-$(date +"%Y-%m-%dT%H-%M-%S").log"

main "$@" 2>&1 | tee $LOGFILE