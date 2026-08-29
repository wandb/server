LSB_DIST=
DIST_VERSION=
PRIVATE_IP=

function discover() {
    discover_lsb
    discover_private_ip

    printf "\n--- System Info ---\n"
    printf "Linux: $LSB_DIST\n"
    printf "Version: $DIST_VERSION\n"
    printf "Private IP: $PRIVATE_IP\n"
    printf "\n"
}

function discover_private_ip() {
    if [ -n "$PRIVATE_ADDRESS" ]; then
        return 0
    fi

    if command_exists "ifconfig"; then
        PRIVATE_IP=$(ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1')
    elif command_exists "ip"; then
        PRIVATE_IP=$(ip -4 addr | grep -oP '(?<=inet\s)\d+(\.\d+){3}' | grep -v '127.0.0.1')
    fi
}


function discover_lsb() {
    if [ -f /etc/os-release ] && [ -r /etc/os-release ]; then
        LSB_DIST="$(. /etc/os-release && echo "$ID")"
        DIST_VERSION="$(. /etc/os-release && echo "$VERSION_ID")"
    else
        bail "Error: Unknown operating system."
    fi
}
