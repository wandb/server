function kubernetes_install_packages() {
    k8sVersion=$1
    folder="${DIR}/packages/kubernetes/${k8sVersion}"
    
    if [[ $AIRGAP != "1" && -n "$AIRGAP" ]]; then
        log_step "Downloading kubelet, kubectl, kubeadm and cni packages"
        pushd $folder
            package_download "cni-plugins.tgz" "https://github.com/containernetworking/plugins/releases/download/v$CNI_PLUGINS_VERSION/cni-plugins-linux-$ARCH-v$CNI_PLUGINS_VERSION.tgz"
            package_download "kubeadm" "https://storage.googleapis.com/kubernetes-release/release/v$k8sVersion/bin/linux/$ARCH/kubeadm"
            package_download "kubelet" "https://storage.googleapis.com/kubernetes-release/release/v$k8sVersion/bin/linux/$ARCH/kubelet"
            package_download "kubectl" "https://storage.googleapis.com/kubernetes-release/release/v$k8sVersion/bin/linux/$ARCH/kubectl"
            package_download "containerd.tar.gz" "https://github.com/containerd/containerd/releases/download/v$k8sVersion/containerd-$k8sVersion-linux-$ARCH.tar.gz"
            package_download "crictl.tar.gz" "https://github.com/kubernetes-sigs/cri-tools/releases/download/v$k8sVersion/crictl-v$k8sVersion-linux-$ARCH.tar.gz"
        popd
    fi

    log_step "Install kubelet, kubectl and cni host packages"
    if kubernetes_host_commands_ok "$k8sVersion"; then
        log_success "Kubernetes host packages already installed"
        return
    fi

    # tar -C /usr/bin -xzf "$folder/assets/crictl-linux-amd64.tar.gz"
    # chmod a+rx /usr/bin/crictl

    # local kubeadmin_path="$folder/$(package_filepath "kubeadm")"
    # cp -f "$kubeadmin_path" /usr/bin/
    # chmod a+rx /usr/bin/kubeadm

    # local kubectl_path="$folder/$(package_filepath "kubectl")"
    # cp -f "$kubectl_path" /usr/bin/
    # chmod a+rx /usr/bin/kubectl

    # local kubelet_path="$folder/$(package_filepath "kubelet")"
    # cp -f "$kubelet_path" /usr/bin/
    # chmod a+rx /usr/bin/kubelet
}

function kubernetes_has_packages() {
    local k8sVersion=$1

    if ! command_exists kubelet; then
        printf "kubelet command missing - will install host components\n"
        return 1
    fi
    if ! command_exists kubeadm; then
        printf "kubeadm command missing - will install host components\n"
        return 1
    fi
    if ! command_exists kubectl; then
        printf "kubectl command missing - will install host components\n"
        return 1
    fi
    if ! ( PATH=$PATH:/usr/local/bin; command_exists kustomize ); then
        printf "kustomize command missing - will install host components\n"
        return 1
    fi
    if ! command_exists crictl; then
        printf "crictl command missing - will install host components\n"
        return 1
    fi
}