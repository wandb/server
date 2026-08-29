#!/bin/bash

function require_root_user() {
    local user="$(id -un 2>/dev/null || true)"
    if [ "$user" != "root" ]; then
        bail "Error: this installer needs to be run as root."
    fi
}

function download_binary() {
    if [ ! -e installer ]; then
        printf "Downloading latest binary"
    fi
}

function download_config() {
    if [ ! -e installer ]; then
        printf "Downloading latest config"
    fi
}

require_root_user
download_config
download_binary

chmod +x installer

./installer install
