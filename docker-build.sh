#!/bin/bash
if [ -z $1 ];then
	echo "Tag não definida.
	Por favor defina $1"
	exit 1
fi
export LC_ALL=C
#GCP
#docker build -t us-east1-docker.pkg.dev/images-registry-410214/docker-images/stocks-helper:$1 . 
#docker push us-east1-docker.pkg.dev/images-registry-410214/docker-images/stocks-helper:$1

#AWS
aws ecr get-login-password --region us-west-1 | sudo docker login --username AWS --password-stdin 905418371571.dkr.ecr.us-west-1.amazonaws.com
sudo docker build -t 905418371571.dkr.ecr.us-west-1.amazonaws.com/stocks-helper:$1 .
sudo docker push 905418371571.dkr.ecr.us-west-1.amazonaws.com/stocks-helper:$1

