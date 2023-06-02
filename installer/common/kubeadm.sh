function kubeadm_api_is_healthy() {
    addr=$PRIVATE_IP:6443
    curl -k "https://$addr/healthz" >/dev/null
}
