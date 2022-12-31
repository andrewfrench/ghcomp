
build: build/inflate build/deflate ## Builds all executables.

build/inflate: ## Builds the inflate executable.
	go build -o out/inflate cmd/inflate/main.go

build/deflate: ## Builds the deflate executable.
	go build -o out/deflate cmd/deflate/main.go

.PHONY: test
test: ## Runs tests.
	go test -race ./...

.PHONY: help
help: ## Prints Makefile help text.
	@grep -E '^[a-zA-Z_\/-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help