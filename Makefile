FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: hello test test-integration fmt
hello:
	echo "Hello World"
test: ## run all unit test
	go test ./... -v -short
test-integration: ## run all integration test
	go test ./... -v -run Integration
fmt: ## format the go source files
	go fmt ./...