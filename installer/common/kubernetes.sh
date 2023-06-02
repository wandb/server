CNI_BIN=/opt/cni/bin

function kubernetes_install_packages() {    
    mkdir -p $PACKAGES
    if [ -z "$AIRGAP" ] || [ "$AIRGAP" != "1" ];  then
        log_step "Downloading packages"
        pushd $PACKAGES > /dev/null 2>&1
            package_download "runc" "https://github.com/opencontainers/runc/releases/download/v$RUNC_VERSION/runc.$ARCH"
            package_download "cni-plugins.tgz" "https://github.com/containernetworking/plugins/releases/download/v$CNI_PLUGINS_VERSION/cni-plugins-linux-$ARCH-v$CNI_PLUGINS_VERSION.tgz"
            package_download "kubeadm" "https://dl.k8s.io/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubeadm"
            package_download "kubelet" "https://dl.k8s.io/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubelet"
            package_download "kubectl" "https://dl.k8s.io/release/v$KUBERNETES_VERSION/bin/linux/$ARCH/kubectl"
            package_download "containerd.tar.gz" "https://github.com/containerd/containerd/releases/download/v$CONTAINERD_VERSION/containerd-$CONTAINERD_VERSION-linux-$ARCH.tar.gz"
            package_download "crictl.tar.gz" "https://github.com/kubernetes-sigs/cri-tools/releases/download/v$CRICTL_VERSION/crictl-v$CRICTL_VERSION-linux-$ARCH.tar.gz"
        popd > /dev/null 2>&1
    fi

    log_step "Installing packages"
    pushd $PACKAGES > /dev/null 2>&1
        printf "Installing containerd\n"
        tar -C /usr/local -xzf "$(package_filepath "containerd.tar.gz")"
        chmod a+rx /usr/bin/kubectl
        kubernetes_configure_containerd_systemd

        printf "Installing runc\n"
        install -m 755 $(package_filepath "runc") /usr/local/sbin/runc

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
        cp -f "$(package_filepath "kubectl")" /usr/bin/
        chmod a+rx /usr/bin/kubectl
    
    
        printf "Installing kubelet\n"
        cp -f "$(package_filepath "kubelet")" /usr/bin/
        chmod a+rx /usr/bin/kubelet
        kubernetes_configure_kubelet_systemd
    popd > /dev/null 2>&1

    printf "Loading Kubelet\n"
    systemctl daemon-reload
    systemctl enable kubelet && systemctl restart kubelet

    log_success "Kubernetes packages installed"
}

function kubernetes_configure_containerd_systemd() {
    mkdir -p /usr/local/lib/systemd/system
    cat > "containerd.service" <<EOF
[Unit]
Description=containerd container runtime
Documentation=https://containerd.io
After=network.target local-fs.target

[Service]
#uncomment to enable the experimental sbservice (sandboxed) version of containerd/cri integration
#Environment="ENABLE_CRI_SANDBOXES=sandboxed"
ExecStartPre=-/sbin/modprobe overlay
ExecStart=/usr/local/bin/containerd

Type=notify
Delegate=yes
KillMode=process
Restart=always
RestartSec=5
# Having non-zero Limit*s causes performance problems due to accounting overhead
# in the kernel. We recommend using cgroups to do container-local accounting.
LimitNPROC=infinity
LimitCORE=infinity
LimitNOFILE=infinity
# Comment TasksMax if your systemd version does not supports it.
# Only systemd 226 and above support this version.
TasksMax=infinity
OOMScoreAdjust=-999

[Install]
WantedBy=multi-user.target
EOF
    cp -f "containerd.service" "/etc/systemd/system/containerd.service"
    chmod 600 /etc/systemd/system/containerd.service

    systemctl daemon-reload
    systemctl enable containerd && systemctl restart containerd
}

function kubernetes_configure_kubelet_systemd() {
    # https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
    
    mkdir -p /etc/systemd/system/kubelet.service.d

    cat > "kubelet.service" <<EOF
[Unit]
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

    cp -f "10-kubeadm.conf" /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
    chmod 600 /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

    systemctl daemon-reload
    systemctl enable kubelet && systemctl restart kubelet
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
    if ! command_exists crictl; then
        printf "crictl command missing - will install host components\n"
        return 1
    fi
}