function kubeadm_api_is_healthy() {
    addr=$PRIVATE_IP:6443
    curl --globoff --noproxy "*" --fail --silent --insecure "https://$addr/healthz" >/dev/null
}

function kubeadm_get_containerd_sock() {
    echo "/run/containerd/containerd.sock"
}