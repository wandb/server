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

function must_swap_off() {
    if swap_is_on || swap_is_enabled; then
        printf "\n${YELLOW}This application is incompatible with memory swapping enabled. Disable swap to continue?${NC} "
        if confirmY ; then
            printf "=> Running swapoff --all\n"
            swapoff --all
            if swap_fstab_enabled; then
                swap_fstab_disable
            fi
            if swap_service_enabled; then
                swap_service_disable
            fi
            if swap_azure_linux_agent_enabled; then
                swap_azure_linux_agent_disable
            fi
            logSuccess "Swap disabled.\n"
        else
            bail "\nDisable swap with swapoff --all and remove all swap entries from /etc/fstab before re-running this script"
        fi
    fi
}

function swap_is_on() {
   swapon --summary | grep --quiet " "
}

function swap_is_enabled() {
    swap_fstab_enabled || swap_service_enabled || swap_azure_linux_agent_enabled
}

function swap_fstab_enabled() {
    cat /etc/fstab | grep --quiet --ignore-case --extended-regexp '^[^#]+swap'
}

function swap_fstab_disable() {
    printf "=> Commenting swap entries in /etc/fstab \n"
    sed --in-place=.bak '/\bswap\b/ s/^/#/' /etc/fstab
    printf "=> A backup of /etc/fstab has been made at /etc/fstab.bak\n\n"
    printf "\n${YELLOW}Changes have been made to /etc/fstab. We recommend reviewing them after completing this installation to ensure mounts are correctly configured.${NC}\n\n"
    sleep 5 # for emphasis of the above ^
}

# This is a service on some Azure VMs that just enables swap
function swap_service_enabled() {
    systemctl -q is-enabled temp-disk-swapfile 2>/dev/null
}

function swap_service_disable() {
    printf "=> Disabling temp-disk-swapfile service\n"
    systemctl disable temp-disk-swapfile
}

function swap_azure_linux_agent_enabled() {
    cat /etc/waagent.conf 2>/dev/null | grep -q 'ResourceDisk.EnableSwap=y'
}

# retry a command if it fails up to $1 number of times
# Usage: cmd_retry 3 curl --globoff --noproxy "*" --fail --silent --insecure https://10.128.0.25:6443/healthz
function cmd_retry() {
    local retries=$1
    shift

    local count=0
    until "$@"; do
        exit=$?
        wait=$((2 ** $count))
        count=$(($count + 1))
        if [ $count -lt $retries ]; then
            echo "Retry $count/$retries exited $exit, retrying in $wait seconds..."
            sleep $wait
        else
            echo "Retry $count/$retries exited $exit, no more retries left."
            return $exit
        fi
    done
    return 0
}

function spinner_until() {
    local timeoutSeconds="$1"
    local cmd="$2"
    local args=${@:3}

    if [ -z "$timeoutSeconds" ]; then
        timeoutSeconds=-1
    fi

    local delay=1
    local elapsed=0
    local spinstr='|/-\'

    while ! $cmd $args; do
        elapsed=$((elapsed + delay))
        if [ "$timeoutSeconds" -ge 0 ] && [ "$elapsed" -gt "$timeoutSeconds" ]; then
            return 1
        fi
        local temp=${spinstr#?}
        printf " [%c]  " "$spinstr"
        local spinstr=$temp${spinstr%"$temp"}
        sleep $delay
        printf "\b\b\b\b\b\b"
    done
}