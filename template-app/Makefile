SVC_NAME := {{.ServiceName}}
APP_PORT ?= 80

GOFILES := ./...
COVERAGE_THRESHOLD := 100

TEST_DOCKER_CONTAINER := test

.PHONY: all
all: build test teardown

.PHONY: setup
setup:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(HOME)/bin latest

.PHONY: build
build:
	docker build -f ./build/Dockerfile --build-arg APP_PORT=$(APP_PORT) -t $(SVC_NAME) .

.PHONY: test
test:
	COVERAGE_THRESHOLD=$(COVERAGE_THRESHOLD) docker-compose -f docker-compose.yaml up --exit-code-from $(TEST_DOCKER_CONTAINER)

.PHONY: teardown
teardown:
	docker-compose -f docker-compose.yaml stop
	docker-compose -f docker-compose.yaml rm --force

.PHONY: fix
fix:
	$(HOME)/bin/golangci-lint run --fix

.PHONY: lint
lint:
	$(HOME)/bin/golangci-lint run
