GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(GOPATH)
VERSION=$(shell git describe --tags --always)
SERVICE=$(shell basename $(CURDIR))
PRIVATE_KEY_NAME=storage/keys/jwt-private.key
PUBLIC_KEY_NAME=storage/keys/jwt-public.key

.PHONY: build-docker
# build docker image
build-docker:
	docker build --build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
		. -t tanyudii/$(SERVICE):latest

.PHONY: build-docker-tag-version
# build docker image with git tag version
build-docker-tag-version:
	docker build --build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
		. -t tanyudii/$(SERVICE):$(VERSION)

.PHONY: oauth-key
# generate oauth key to env file
oauth-key:
	openssl genrsa -out $(PRIVATE_KEY_NAME) 4096 && \
	openssl rsa -in $(PRIVATE_KEY_NAME) -pubout -out $(PUBLIC_KEY_NAME)

help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help