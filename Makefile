NAME := maze
TARGET := ./cmd/$(NAME)
VERSION := v0.0.1
ifeq ($(COMMIT),)
COMMIT := $(shell git rev-parse HEAD)
endif
DATE := $(shell date "+%Y-%m-%dT%H:%M:%S%z")

SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.version=$(VERSION)\" -X \"main.commit=$(COMMIT)\" -X \"main.date=$(DATE)\" -extldflags -static"
NOVENDOR := $(shell go list ./... | grep -v vendor)

export GO111MODULE=on

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(SRCS)
	CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(NAME) $(TARGET)

.PHONY: ci-test
ci-test:
	go test -coverpkg=./... -coverprofile=coverage.txt -v ./...

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf dist/*
	rm -rf vendor/*

.PHONY: fast
fast:
	CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(NAME) $(TARGET)

.PHONY: cross-build
cross-build:
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME) $(TARGET); \
		done; \
	done

dist/linux-amd64/$(NAME): cross-build

DOCKER_REPOSITORY := docker.io
ifeq ($(DOCKER_REPOSITORY),)
DOCKER_IMAGE_NAME := masumi/maze
else
DOCKER_IMAGE_NAME := $(DOCKER_REPOSITORY)/masumi/maze
endif
DOCKER_IMAGE_TAG  ?= latest
DOCKER_IMAGE      := $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

.PHONY: docker-build
docker-build:
	docker build --build-arg COMMIT=$(COMMIT) -t $(DOCKER_IMAGE) .

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: test
test:
	go test -cover -v $(NOVENDOR)
