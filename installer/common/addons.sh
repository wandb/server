function addons_download() {
    mkdir -p $MANIFEST
    if [ -z "$AIRGAP" ] || [ "$AIRGAP" != "1" ];  then
        log_step "Downloading addons"

        local manifest_path=$MANIFEST/contour.yaml
        # TODO: figure out a way to lock the script down
        local manifest_url="https://projectcontour.io/quickstart/contour.yaml"
        package_download_url_with_retry "$manifest_url"  "$manifest_path"
    fi
}

function addons_install() {
    mkdir -p $MANIFEST

    local contour=$MANIFEST/contour.yaml
    kubectl apply -f $contour
}
