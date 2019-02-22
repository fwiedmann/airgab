#!/usr/bin/env bash

set -e

if [[ -z "$DEBUG" ]]; then
    set -x
fi

# check if ssh key is mounted
if [[ ! -f "/home/pilot/.ssh/id_rsa"  ]];then
    echo 'No private ssh-key was found. Please mount the key like this: "/home/pilot/.ssh/id_rsa"'
    exit 1
fi

# change owner of ssh key to container user
chown pilot:pilot /home/pilot/.ssh/id_rsa

# run airgab
./airgab "$@"
