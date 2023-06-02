CNI_BIN=/opt/cni/bin

function kubernetes_install_packages() {    
    mkdir -p $PACKAGES
    if [ -z "$AIRGAP" ] || [ "$AIRGAP" != "1" ];  then
        log_step "Downloading kubelet, kubectl, kubeadm and cni packages"
        pushd $PACKAGES > /dev/null 2>&1
            package_download "cni-plugins.tgz" "https://github.com/containernetworking/plugins/releases/download/v$CNI_PLUGINS_VERSION/cni-plugins-linux-$ARCH-v$CNI_PLUGINS_VERSION.tgz"
            package_download "kubeadm" "https://dl.k8s.io/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubeadm"
            package_download "kubelet" "https://dl.k8s.io/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubelet"
            package_download "kubectl" "https://dl.k8s.io/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubectl"
            package_download "containerd.tar.gz" "https://github.com/containerd/containerd/releases/download/v$CONTAINERD_VERSION/containerd-$CONTAINERD_VERSION-linux-$ARCH.tar.gz"
            package_download "crictl.tar.gz" "https://github.com/kubernetes-sigs/cri-tools/releases/download/v$CRICTL_VERSION/crictl-v$CRICTL_VERSION-linux-$ARCH.tar.gz"
        popd > /dev/null 2>&1
    fi

    # log_step "Install kubelet, kubectl and cni host packages"
    # if kubernetes_host_commands_ok "$k8sVersion"; then
    #     log_success "Kubernetes host packages already installed"
    #     return
    # fi

    log_step "Installing packages"
    pushd $PACKAGES > /dev/null 2>&1
        log_step "Installing cni plugins\n"
        mkdir -p $CNI_BIN
        tar -C $CNI_BIN -xzf "$(package_filepath "cni-plugins.tgz")"

        printf "Installing crictl\n"
        tar -C /usr/bin -xzf $(package_filepath "crictl.tar.gz")
        chmod a+rx /usr/bin/crictl

        printf "Installing kubeadm\n"
        cp -f "$(package_filepath "kubeadm")" /usr/bin/
        chmod a+rx /usr/bin/kubeadm

        printf "Installing kubectl\n"
        # cp -f "$(package_filepath "kubectl")" /usr/bin/
        # chmod a+rx /usr/bin/kubectl
    
        printf "Installing kubelet\n"
        kubernetes_configure_kubelet_systemd
        cp -f "$(package_filepath "kubelet")" /usr/bin/
        chmod a+rx /usr/bin/kubelet
    popd > /dev/null 2>&1

    printf "Restarting Kubelet"
    systemctl daemon-reload
    systemctl enable kubelet && systemctl restart kubelet

    log_success "Kubernetes packages installed"
}

function kubernetes_configure_kubelet_systemd() {
    # https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
    
    mkdir -p /etc/systemd/system/kubelet.service.d

    cat > "kubelet.service" <<EOF [Unit]
Description=kubelet: The Kubernetes Node Agent
Documentation=https://kubernetes.io/docs/home/
Wants=network-online.target
After=network-online.target

[Service]
ExecStart=/kubelet
Restart=always
StartLimitInterval=0
RestartSec=10

[Install]
WantedBy=multi-user.target"
EOF
    cp -f "kubelet.service" "/etc/systemd/system/kubelet.service"
    chmod 600 /etc/systemd/system/kubelet.service

    cat > "10-kubeadm.conf" <<EOF
# Note: This dropin only works with kubeadm and kubelet v1.11+
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_CONFIG_ARGS=--config=/var/lib/kubelet/config.yaml"
# This is a file that "kubeadm init" and "kubeadm join" generates at runtime, populating the KUBELET_KUBEADM_ARGS variable dynamically
EnvironmentFile=-/var/lib/kubelet/kubeadm-flags.env
# This is a file that the user can use for overrides of the kubelet args as a last resort. Preferably, the user should use
# the .NodeRegistration.KubeletExtraArgs object in the configuration files instead. KUBELET_EXTRA_ARGS should be sourced from this file.
EnvironmentFile=-/etc/default/kubelet
ExecStart=
ExecStart=/usr/bin/kubelet \$KUBELET_KUBECONFIG_ARGS \$KUBELET_CONFIG_ARGS \$KUBELET_KUBEADM_ARGS $KUBELET_EXTRA_ARGS
EOF

    mkdir -p /etc/systemd/system/kubelet.service.d
    cp -f "10-kubeadm.conf" /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
    chmod 600 /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

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