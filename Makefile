IMAGE_TAG = 1.0

SSH_FOLDER = $$HOME


all:  prepare build run

prepare:
	mkdir -pv ./backup

cleanup:
	sudo rm -rf ./backup
	docker rmi --force airgab wiedmannfelix/airgab:$(IMAGE_TAG)  wiedmannfelix/airgab:latest

build:
	docker build -t airgab .

run: 
	docker run -it -v $$PWD/backup:/home/pilot/backup -v $(SSH_FOLDER)/.ssh/id_rsa:/home/pilot/.id_rsa -p 9100:9100 airgab --user=pi --host=192.168.2.233 --source=/opt/ghost --destination=/home/pilot/backup/ --options=-a --interval=10s --private-key=$$HOME/.ssh/id_rsa

push:
	docker tag airgab wiedmannfelix/airgab:$(IMAGE_TAG)
	docker tag airgab wiedmannfelix/airgab:latest
	docker push wiedmannfelix/airgab:$(IMAGE_TAG)
	docker push wiedmannfelix/airgab:latest
