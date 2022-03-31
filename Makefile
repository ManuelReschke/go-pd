FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: hello
hello:
	echo "Hello World"
.PHONY: hello

test: ## run all unit test
	go test ./... -v -short
.PHONY: test

test-integration: ## run all integration test
	go test ./... -v -run Integration
.PHONY: test-integration

fmt: ## format the go source files
	go fmt ./...
.PHONY: fmt

build:
	go build -o bin/go-pd
.PHONY: build