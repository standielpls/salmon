.PHONY: all

test: ## Run all the tests
	echo 'mode: atomic' > coverage.txt && go test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

ci: lint test ## Run all the tests and code checks

build: ## Build a version
	go build -o ./bin/salmon ./cmd/main.go 

run: build ## Build & Run 
	./bin/salmon

install: 
	go build -i -o ./bin/salmon ./cmd/main.go

clean: ## Remove temporary files
	go clean

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build