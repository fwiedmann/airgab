#!/usr/bin/env bash

set -e

if [[ -z "$DEBUG" ]]; then
    set -x
fi

# check if ssh key is mounted
if [[ ! -f "/home/pilot/.id_rsa"  ]];then
    echo 'No private ssh-key was found. Please mount the key like this: "/home/pilot/.ssh/id_rsa"'
    exit 1
fi

# copy ssh-key to users ssh-fodler to prevent issues with ownership
cp .id_rsa .ssh/id_rsa

# change owner of ssh key to container user
chown pilot:pilot /home/pilot/.ssh/id_rsa

# step down from root and  run airgab
su pilot -c "./airgab $*"