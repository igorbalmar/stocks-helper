#!/bin/bash
if [ -z $1 ]
then
	echo "Tag não definida.
	Por favor defina $1"
	exit 1
fi
export LC_ALL=C
sudo docker build -t us-east1-docker.pkg.dev/images-registry-410214/docker-images/stocks-helper:$1 . && \
	sudo docker push us-east1-docker.pkg.dev/images-registry-410214/docker-images/stocks-helper:$1
