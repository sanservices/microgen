.DEFAULT_GOAL := help
BINARY := {{ cookiecutter.root_directory_name }}

.PHONY: help run build test tidy fmt lint generate {% if cookiecutter.use_mcp == 'y' %}mcp-client {% endif %}docker-build docker-run

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

{% if cookiecutter.use_mcp == 'y' %}
mcp-client: ## Run the example MCP client (service must be running)
	go run ./internal/mcp/client
{% endif %}
docker-build: ## Build the docker image
	docker build -t $(BINARY):latest .

docker-run: ## Run the docker image (REST + gRPC{% if cookiecutter.use_mcp == 'y' %} + MCP{% endif %})
	docker run --rm -p 8080:8080 -p 50051:50051 {% if cookiecutter.use_mcp == 'y' %}-p 8081:8081 {% endif %}$(BINARY):latest
