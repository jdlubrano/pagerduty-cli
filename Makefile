.DEFAULT_GOAL := build

goreleaser := $(shell command -v goreleaser 2> /dev/null)

build: test
	go build -v -a

test: ## Run all tests
	go test ./... -v

install_goreleaser:
ifdef goreleaser
else
	$(error You need to install goreleaser; see https://goreleaser.com/install)
endif

define_github_token:
ifeq ($(GITHUB_TOKEN),)
	$(error You need to define a GITHUB_TOKEN with a "repo" scope)
endif

test_release: define_github_token install_goreleaser build
	@goreleaser release --skip-publish --rm-dist

release: test define_github_token install_goreleaser
	@goreleaser release --rm-dist

.PHONY: test build setup install_dep help
