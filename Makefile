VERSION ?= latest
BUILDFLAGS := -ldflags="-s -w" -gcflags=-trimpath=$(CURDIR)
IMAGE_NAME := proxytest
IMAGE_REGISTRY ?= ghcr.io/razzie
FULL_IMAGE_NAME := $(IMAGE_REGISTRY)/$(IMAGE_NAME):$(VERSION)

.PHONY: build
build:
	go build $(BUILDFLAGS) .

.PHONY: docker-build
docker-build:
	docker build . -t $(FULL_IMAGE_NAME)

.PHONY: docker-push
docker-push: docker-build
	docker push $(FULL_IMAGE_NAME)

.PHONY: deploy
deploy:
	helm update --install proxytest helm-chart/proxytest/

.PHONY: undeploy
undeploy:
	helm uninstall proxytest
