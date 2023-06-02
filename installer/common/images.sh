function image_download() {
    local name=$1
    local url=$2

    printf "Downloading $name\n"

    if [ -f "$name" ]; then
        printf "$name already downloaded\n"
        return
    fi

    if command_exists docker; then
        docker pull $url > /dev/null 2>&1
        docker save $url | gzip > $name.tar.gz
    elif command_exists ctr; then
        ctr images pull --plain-http $url > /dev/null 2>&1
        ctr --namespace=default images export - $url | gzip > $name.tar.gz
    else
        log_warning "No support client installed for pulling images"
    fi
}

function images_download() {
    log_step "Downloading images"

    mkdir -p $IMAGES
    if [ -z "$AIRGAP" ] || [ "$AIRGAP" != "1" ];  then
        pushd $IMAGES > /dev/null 2>&1
            image_download "coredns" $IMAGE_COREDNS
            image_download "etcd" $IMAGE_ETCD
            image_download "kube-apiserver" $IMAGE_KUBE_API
            image_download "kube-controller-manager" $IMAGE_KUBE_CONTROLLER
            image_download "kube-scheduler" $IMAGE_KUBE_SCHEDULER
            image_download "kube-proxy" $IMAGE_KUBE_PROXY
            image_download "pause" $IMAGE_PAUSE
        popd > /dev/null 2>&1
    fi
}

function images_load() {
    find "$1" -type f | xargs -I {} bash -c "cat {} | gunzip | ctr -a $(kubeadm_get_containerd_sock) -n=k8s.io images import -"
}
