#!/usr/bin/env bash

set -e
set -o pipefail


SOURCE_DATA_DIRCECTORY="$1"
DESTIANTION_DATA_DIRECTORY="$2"
RSYNC_OPTIONS=${3:-"-a"}

function check() {
    if [[ ! "$(type rsync)" ]]; then
     echo "rsync is not installed on your system. Please install it via your packagemanager, for example: apt install rsync"
     exit 1
    fi

    if [[ -z "$SOURCE_DATA_DIRCECTORY" ]]; then
        echo '$SOURCE_DATA_DIRCECTORY is not set. This will point to the directory which will be backuped'
        exit 1
    fi

    if [[ -z "$DESTIANTION_DATA_DIRECTORY" ]]; then
        echo '$DESTIANTION_DATA_DIRECTORY is not set. It will point to the directory where your data will be backuped'
        exit 1
    fi
}

function backup() {
    rsync "$RSYNC_OPTIONS" "$SOURCE_DATA_DIRCECTORY" "$DESTIANTION_DATA_DIRECTORY"
}


check
backup
