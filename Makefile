dummy := $(shell touch .env)
include .env
export

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Helpers

.PHONY: check
check: %: format tidy lint test ## Run format, tidy, lint, and test commands

.PHONY: test
test: ## Run tests
	go test -v ./...

.PHONY: lint
lint: ## Run golang linters
	golangci-lint run

.PHONY: tidy
tidy: ## Tidy up the go.mod dependencies
	go mod tidy -v

PHONY: format
format: ## Run go format
	go fmt ./...

##@ Run apps

.PHONY: run
run: ## Run the application in Docker
	go run ./cmd/
