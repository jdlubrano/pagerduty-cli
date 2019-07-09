.DEFAULT_GOAL := build

build: test
	go build -v -a

test: ## Run all tests
	go test ./... -v

.PHONY: test build setup install_dep help
