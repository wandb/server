function install_dependencies() {
    dependencies_download
}

function dependencies_download() {
    mkdir -p $DEPENDENCIES
    if [ -z "$AIRGAP" ] || [ "$AIRGAP" != "1" ];  then
        log_step "Downloading host dependencies"
        pushd $DEPENDENCIES > /dev/null 2>&1
            package_download "openssl.tar.gz" "https://www.openssl.org/source/openssl-$OPENSSL_VERSION.tar.gz"
        popd > /dev/null 2>&1
    fi
}
