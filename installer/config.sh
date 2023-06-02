export AIRGAP=0
export ARCH="amd64"

export KUBERNETES_VERSION="1.27.2"
export CRICTL_VERSION="1.27.0"
export CNI_PLUGINS_VERSION="1.3.0"
export CONTAINERD_VERSION="1.7.1"
export RUNC_VERSION="1.1.7"

export OPENSSL_VERSION="3.1.1"

export IMAGE_COREDNS=registry.k8s.io/coredns/coredns:v1.10.1
export IMAGE_ETCD=registry.k8s.io/etcd:3.5.7-0
export IMAGE_KUBE_API=registry.k8s.io/kube-apiserver:v$KUBERNETES_VERSION
export IMAGE_KUBE_CONTROLLER=registry.k8s.io/kube-controller-manager:v$KUBERNETES_VERSION
export IMAGE_KUBE_PROXY=registry.k8s.io/kube-proxy:v$KUBERNETES_VERSION
export IMAGE_KUBE_SCHEDULER=registry.k8s.io/kube-scheduler:v$KUBERNETES_VERSION
export IMAGE_PAUSE=registry.k8s.io/pause:3.9

export PACKAGES=$DIR/packages/kubernetes/$KUBERNETES_VERSION
export IMAGES=$DIR/packages/kubernetes/$KUBERNETES_VERSION/images
export DEPENDENCIES=$DIR/packages/deps
export HOSTNAME="$(hostname | tr '[:upper:]' '[:lower:]')"