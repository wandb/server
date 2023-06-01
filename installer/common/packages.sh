function package_filepath() {
    local package="$1"
    mkdir -p assets
    echo "assets/${package}"
}

function package_download_url_with_retry() {
    local url="$1"
    local filepath="$2"
    local max_retries="${3:-10}"

    local errcode=
    local i=0
    while [ $i -ne "$max_retries" ]; do
        errcode=0
        curl -fL -o "${filepath}" "${url}" || errcode="$?"
        # 18 transfer closed with outstanding read data remaining
        # 56 recv failure (connection reset by peer)
        if [ "$errcode" -eq "18" ] || [ "$errcode" -eq "56" ]; then
            i=$(($i+1))
            continue
        fi
        return "$errcode"
    done
    return "$errcode"
}

function package_download() {
    local package="$1"
    local package_url="$2"

    if [ -z "$package" ]; then
        bail "package_download called with no package name"
    fi

    local etag="$(grep -F "${package}" assets/Manifest | awk 'NR == 1 {print $2}')"
    local checksum="$(grep -F "${package}" assets/Manifest | awk 'NR == 1 {print $3}')"

    local newetag="$(curl -IfsSL "$package_url" | grep -i 'etag:' | sed -r 's/.*"(.*)".*/\1/')"
    if [ -n "${etag}" ] && [ "${etag}" = "${newetag}" ]; then
        echo "Package ${package} already exists, not downloading"
        return
    fi

    log_step "Downloading ${package}"

    local filepath="$(package_filepath "${package}")"
    package_download_url_with_retry "$package_url" "$filepath"

    checksum="$(md5sum "${filepath}" | awk '{print $1}')"
    echo "${package} ${newetag} ${checksum}" >> assets/Manifest
}

