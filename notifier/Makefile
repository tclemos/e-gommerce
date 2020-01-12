SERVICE_NAME=notifier

DOCKER_NAME="e-gommerce-${SERVICE_NAME}"
DOCKER_NETWORK=e-gommerce
DOCKER_INTERNAL_PORT=8080
DOCKER_EXTERNAL_PORT=8080
DOCKER_WORKDIR=/src
DOCKER_IMG=golang:1.13.4-alpine

.PHONY: protocol
protocol:
	protoc --go_out=plugins=grpc:. ./pkg/protocol/*.proto

.PHONY: build
build:
	CGO_ENABLED=0 \
	GO111MODULE=on \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o ./build/${SERVICE_NAME} ./cmd/${SERVICE_NAME}/main.go

.PHONY: debug
debug:
	$(MAKE) DEBUG=-debug run

.PHONY: run
run: build
	sudo docker run --rm --name ${DOCKER_NAME} \
			   --network ${DOCKER_NETWORK} \
			   -v $$(pwd):${DOCKER_WORKDIR} \
			   -w ${DOCKER_WORKDIR} \
			   -p ${DOCKER_EXTERNAL_PORT}:${DOCKER_INTERNAL_PORT} \
			   -it \
			   ${DOCKER_IMG} \
			   ./build/${SERVICE_NAME} ${DEBUG}
