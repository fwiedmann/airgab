all: build run

build:
	go build

run:
	./airgab --user=pi --host=192.168.2.233 --source=/opt/ghost --destination=/home/fwiedmann --options=-a --interval=10s --private-key=/home/fwiedmann/.ssh/id_rsa