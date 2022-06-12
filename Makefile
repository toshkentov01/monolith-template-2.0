CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

IMG_NAME=${APP}
REGISTRY=${REGISTRY:-861701250313.dkr.ecr.us-east-1.amazonaws.com}
TAG=latest
ENV_TAG=latest

ifneq (,$(wildcard ./.env))
	include .env
endif

ifdef CI_COMMIT_BRANCH
        include .build_info
endif

make create-env:
	cp ./.env.example ./.env

set-env:
	./scripts/set-env.sh ${CURRENT_DIR}

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

lint:
	golint -set_exit_status ${PKG_LIST}

unit-tests: set-env ## Run unit-tests
	go test -mod=vendor -v -cover -short ${PKG_LIST}

race: set-env ## Run data race detector
	go test -mod=vendor -race -short ${PKG_LIST}

msan: set-env ## Run memory sanitizer. If this test fails, you need to write the following command: export CC=clang (if you have installed clang)
	env CC=clang env CXX=clang++ go test -mod=vendor -msan -short ${PKG_LIST}

delete-branches:
	${CURRENT_DIR}/scripts/delete-branches.sh

swag-gen:
	echo ${REGISTRY}
	swag init -g cmd/main.go -o api/docs


ifneq (,$(wildcard vendor))
	go mod vendor
endif

.PHONY: vendor
