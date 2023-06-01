#!/bin/bash

set -e

DIR=.

# <Config>
export AIRGAP=
export ARCH="amd64"

export KUBERNETES_VERSION="1.27.2"
export CRICTL_VERSION="1.27.0"
export CNI_PLUGINS_VERSION="1.3.0"
export CONTAINERD_VERSION="1.7.1"
# </Config>

# <ImportInline>
. $DIR/installer/common/kubernetes.sh
. $DIR/installer/common/logging.sh
. $DIR/installer/common/packages.sh
. $DIR/installer/common/semver.sh
. $DIR/installer/common/utils.sh
# </ImportInline>

CNI_DEST=/opt/cni/bin
KUBECONFIG=/etc/kubernetes/admin.conf

function main() {
    log_step "Running install with the argument(s): $*"

    require_root_user
    path_add "/usr/local/bin"

    if [ "$AIRGAP" = "1" ]; then
        log_step "Running in airgapped enviroment."
    fi

    export KUBECONFIG=/etc/kubernetes/admin.conf
    folder="$DIR/packages/kubernetes/$KUBERNETES_VERSION"
    
    mkdir -p $folder
    pushd $folder
        package_download "cni-plugins.tgz" "https://github.com/containernetworking/plugins/releases/download/v$CNI_PLUGINS_VERSION/cni-plugins-linux-$ARCH-v$CNI_PLUGINS_VERSION.tgz"
        package_download "kubeadm" "https://storage.googleapis.com/kubernetes-release/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubeadm"
        package_download "kubelet" "https://storage.googleapis.com/kubernetes-release/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubelet"
        package_download "kubectl" "https://storage.googleapis.com/kubernetes-release/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubectl"
        package_download "containerd.tar.gz" "https://github.com/containerd/containerd/releases/download/v$CONTAINERD_VERSION/containerd-$CONTAINERD_VERSION-linux-$ARCH.tar.gz"
        package_download "crictl.tar.gz" "https://github.com/kubernetes-sigs/cri-tools/releases/download/v$KUBERNETES_VERSION/crictl-v$KUBERNETES_VERSION-linux-amd64.tar.gz"
    popd

    # mkdir -p $CNI_DEST
    # tar xf "assets/cni-plugins.tgz" -C $CNI_DEST

    printf "\n"
}

LOGS_DIR="$DIR/logs"
mkdir -p $LOGS_DIR
LOGFILE="$LOGS_DIR/install-$(date +"%Y-%m-%dT%H-%M-%S").log"

main "$@" 2>&1 | tee $LOGFILE