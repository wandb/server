function common_list_images_in_manifest_file() {
    local manifest_file="$1"

    local image_list=
    for image in $(grep "^image " "$manifest_file" | awk '{print $3}' | tr '\n' ' ') ; do
        image_list=$image_list" $(canonical_image_name "$image")"
    done
    echo "$image_list" | xargs # trim whitespace
}

function asset_download_url_in_manifest() {
    local manifest_file="$1"
    local package="$2"

    local download_url=$(grep "^asset $package" "$manifest_file" | awk '{print $3}' | tr '\n' ' ')

    echo "$download_url" | xargs # trim whitespace
}