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
    images_download
    images_load $IMAGES
}

function init() {
    set +o pipefail

    cmd_retry 3 kubeadm init \
        --ignore-preflight-errors="all" \
        | tee /tmp/kubeadm-init
    
    log_step "Waiting for kubernetes api health to report ok"
    if ! spinner_until 120 kubeadm_api_is_healthy; then
        bail "Kubernetes API failed to report healthy"
    fi

    printf "\n\n"
    
    export KUBECONFIG=/etc/kubernetes/admin.conf
    
    kubectl cluster-info
    log_success "Cluster initialized"
}

# at this point kubectl should be configured.
function addons() {
    addons_download
    addons_install
}

function main() {
    log_step "Running install with the argument(s): $*"
    
    discover
    setup  
    init
    
    printf "\n"
}

LOGS_DIR="$DIR/logs"
mkdir -p $LOGS_DIR
LOGFILE="$LOGS_DIR/install-$(date +"%Y-%m-%dT%H-%M-%S").log"

main "$@" 2>&1 | tee $LOGFILE