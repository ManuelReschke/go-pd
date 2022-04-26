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

coverage: ## create coverage report with go get golang.org/x/tools/cmd/cover
	go test -cover -coverprofile=c.out ./...
	go tool cover -html=c.out -o coverage.html
.PHONY: coverage

fmt: ## format the go source files
	go fmt ./...
.PHONY: fmt

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/go-pd
	env GOOS=linux GOARCH=arm64 go build -o bin/arm/go-pd
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/go-pd.exe
.PHONY: build