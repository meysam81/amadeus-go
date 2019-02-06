.PHONY: proto

IMAGE_NAME="amadeus-go"
NETWORK="amadeus"
SERVICE_DIR="/go/src/amadeus-go"
PWD=$(shell pwd)

proto:
	protoc --go_out=plugins=grpc:${PWD} pb/amadeus/amadeus.proto

build: proto
	docker build -t ${IMAGE_NAME} .


up:
	docker network ls | grep ${NETWORK} | [ `wc -l` -ne 1 ] && \
		docker network create ${NETWORK} || echo "\`${NETWORK}\` network already exists"
	docker run -v ${PWD}:${SERVICE_DIR}:ro \
		--net ${NETWORK} \
		--name ${IMAGE_NAME} \
		-dp 8000:8000 \
		${IMAGE_NAME}

down:
	docker container stop ${IMAGE_NAME} | \
		xargs docker rm
