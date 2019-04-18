IMAGE_TAG = 1.0

SSH_FOLDER = $$HOME


all:  prepare build run

prepare:
	mkdir -pv ./backup

cleanup:
	sudo rm -rf ./backup

build:
	GO111MODULE=on CGO_ENABLED=0 go build -a -o airgab .
