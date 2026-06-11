.DEFAULT_GOAL := help
BINARY := {{ cookiecutter.root_directory_name }}

.PHONY: help run build test tidy fmt lint generate docker-build docker-run

help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'

run: ## Run the service locally
	go run .

build: ## Build the binary
	go build -o $(BINARY) .

test: ## Run the test suite
	go test ./...

tidy: ## Sync go.mod/go.sum with the source
	go mod tidy

fmt: ## Format the code
	gofmt -l -s -w .

lint: ## Run golangci-lint
	golangci-lint run

generate: ## Generate protobuf / gRPC / gateway / OpenAPI code
	buf dep update
	buf generate

docker-build: ## Build the docker image
	docker build -t $(BINARY):latest .

docker-run: ## Run the docker image (REST + gRPC)
	docker run --rm -p 8080:8080 -p 50051:50051 $(BINARY):latest
