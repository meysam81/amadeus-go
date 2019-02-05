IMAGE_NAME="amadeus-go"
NETWORK="amadeus"
SERVICE_DIR="/go/src/amadeus-go"
PWD=`pwd`

build:
	docker build -t ${IMAGE_NAME} .

test:
	echo ${PWD}
run:

	docker network ls | grep ${NETWORK} | [ `wc -l` -ne 1 ] && \
		docker network create ${NETWORK} || echo "\`${NETWORK}\` network already exists"
	docker run --rm -v ${PWD}:${SERVICE_DIR}:ro \
		--net ${NETWORK} \
		--name ${IMAGE_NAME} \
		-dp 8000:8000 \
		-e MICRO_SERVER_ADDRESS=:8000 \
		-e MICRO_REGISTRY=mdns \
		${IMAGE_NAME}
