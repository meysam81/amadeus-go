.PHONY: proto

IMAGE_NAME="amadeus-go"
NETWORK="amadeus"
PROJECT_ROOT="/go/src/amadeus-go"
PWD=$(shell pwd)
PROTO_FILES=$(shell grep --exclude=*.pb.go -r .proto$  ${PWD}/api/ | cut -d: -f1 | sort | uniq)
PROTO_IMPORT=$(shell dirname ${PWD})

proto:
	for i in ${PROTO_FILES}; do protoc -I ${PROTO_IMPORT} --go_out=plugins=grpc:${PROTO_IMPORT} $$i; done

build:
	docker build -t ${IMAGE_NAME} .

up:
	docker network ls | grep ${NETWORK} | [ `wc -l` -ne 1 ] && \
		docker network create ${NETWORK} || echo "\`${NETWORK}\` network already exists"
	docker run -v ${PWD}:${PROJECT_ROOT}:ro \
		--net ${NETWORK} \
		--name ${IMAGE_NAME} \
		-dp 8000:8000 \
		${IMAGE_NAME}

down:
	docker container stop ${IMAGE_NAME} | \
		xargs docker rm

dev_run:
	go run ./cmd/srv/srv.go

dev_cli:
	go run ./cmd/cli/cli.go

