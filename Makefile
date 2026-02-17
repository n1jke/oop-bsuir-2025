COVERAGE_FILE ?= coverage.out

TARGET_PKG ?= cmd/bank_app               
BINARY_NAME ?= app                 
WORK_DIR ?= laboratory_work-1

# Build 
.PHONY: build
build:
	@echo "go build  ${WORK_DIR}/${TARGET}"
	@mkdir -p .bin
	@cd $(WORK_DIR) && go build -o ../.bin/$(BINARY_NAME) ./$(TARGET_PKG)

# Test
.PHONY: test
test:
	@cd $(WORK_DIR) && go test -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'

.PHONY: test_race
test_race:
	@cd $(WORK_DIR) && go test --race -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'

.PHONY: html_test
html_test: 
	@go tool cover -html='$(COVERAGE_FILE)' -o coverage.html
	@echo "Coverage report saved to coverage.html"

# Lint
.PHONY: fmt
fmt:
	@echo "cd $(WORK_DIR) && go fmt ./..."
	@cd $(WORK_DIR) &&  go fmt ./...

.PHONY: lint
lint:
	@golangci-lint --version && echo "cd $(WORK_DIR) &&  golangci-lint -v run --fix ./..." || echo "golangci-lint not found"
	@cd $(WORK_DIR) && golangci-lint -v run --fix ./...

# Cleanup
.PHONY: clean
clean: 
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f coverage.out
	@rm -f coverage.html
