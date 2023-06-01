function require_root_user() {
    local user="$(id -un 2>/dev/null || true)"
    if [ "$user" != "root" ]; then
        bail "Error: this installer needs to be run as root."
    fi
}

function path_add() {
    if [ -d "$1" ] && [[ ":$PATH:" != *":$1:"* ]]; then
        PATH="${PATH:+"$PATH:"}$1"
    fi
}

function command_exists() {
    command -v "$@" > /dev/null 2>&1
}

function is_valid_ipv4() {
    if echo "$1" | grep -qs '^[0-9][0-9]*\.[0-9][0-9]*\.[0-9][0-9]*\.[0-9][0-9]*$'; then
        return 0
    else
        return 1
    fi
}

function is_valid_ipv6() {
    if echo "$1" | grep -qs "^\([0-9a-fA-F]\{0,4\}:\)\{1,7\}[0-9a-fA-F]\{0,4\}$"; then
        return 0
    else
        return 1
    fi
}

function rm_file() {
    if [ -f "$1" ]; then
        rm $1
    fi
}
