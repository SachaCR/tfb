.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run the service
	go run main.go

.PHONY: build
build: ## Build the service
	go build 

.PHONY: install
install: ## Install Spanny in your $GOPATH/bin
	go install ./cmd/spanny

.PHONY: test
test: ## Test the service with default test runner
	go test -coverprofile=coverage.out ./...

.PHONY: coverage
coverage: test ## Open test coverage 
	go tool cover -html=coverage.out

.PHONY: check
check: 
	go run main.go Dorian -r D && go run main.go chord 5-5-5-5-5-5 -r A && go run main.go scale C-D-E-F-G-A-B -i U && go run main.go chord 0-2-2-0-0-0 -r E && go run main.go scale C-D-E-F-G-A-B -m note
