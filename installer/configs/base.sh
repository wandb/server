export PACKAGES=$DIR/packages/kubernetes/$KUBERNETES_VERSION
export IMAGES=$DIR/packages/kubernetes/$KUBERNETES_VERSION/images
export MANIFEST=$DIR/packages/kubernetes/$KUBERNETES_VERSION/manifests
export DEPENDENCIES=$DIR/packages/deps
export HOSTNAME="$(hostname | tr '[:upper:]' '[:lower:]')"