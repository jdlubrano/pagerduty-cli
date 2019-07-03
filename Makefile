.DEFAULT_GOAL := build

build: test
	go build

test: ## Run all tests
	go test ./...

.PHONY: test build setup install_dep help
