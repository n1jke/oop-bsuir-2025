COVERAGE_FILE ?= coverage.out

TARGET ?= app ## change to you cmd/ folder name

# Build 
.PHONY: build
build:
	@echo "go build  ${TARGET}"
	@mkdir -p .bin
	@go build -o ./bin/${TARGET} ./cmd/${TARGET}

# Test
.PHONY: test
test:
	@go test -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'

.PHONY: test_race
test_race:
	@go test --race -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'

.PHONY: html_test
html_test: 
	@go tool cover -html='$(COVERAGE_FILE)' -o coverage.html
	@echo "Coverage report saved to coverage.html"

# Lint
.PHONY: fmt
fmt:
	@echo "go fmt ./..."
	@go fmt ./...

.PHONY: lint
lint:
	@golangci-lint --version && echo "golangci-lint -v run --fix ./..." || echo "golangci-lint not found"
	@golangci-lint -v run --fix ./...

# Cleanup
.PHONY: clean
clean: 
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f coverage.out
	@rm -f coverage.html
