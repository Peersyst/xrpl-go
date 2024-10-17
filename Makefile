.PHONY: lint test benchmark

EXCLUDED_TEST_PACKAGES = $(shell go list ./... | grep -v /faucet | grep -v /examples)
EXCLUDED_COVERAGE_PACKAGES = $(shell go list ./... | grep -v /faucet | grep -v /examples | grep -v /testutil | grep -v /interfaces)

PARALLEL_TESTS = 4
TEST_TIMEOUT = 5m

################################################################################
############################### LINTING ########################################
################################################################################

lint:
	@echo "Linting Go code..."
	@golangci-lint run
	@echo "Linting complete!"

lint-fix:
	@echo "Fixing Go code..."
	@gofmt -w -s .
	@echo "Fixing complete!"

################################################################################
############################### TESTING ########################################
################################################################################

test-all:
	@echo "Running Go tests..."
	@go test -v $(EXCLUDED_TEST_PACKAGES)
	@echo "Tests complete!"

test-binary-codec:
	@echo "Running Go tests for binary codec package..."
	@go test -v ./binary-codec/...
	@echo "Tests complete!"

test-address-codec:
	@echo "Running Go tests for address codec package..."
	@go test -v ./address-codec/...
	@echo "Tests complete!"

test-keypairs:
	@echo "Running Go tests for keypairs package..."
	@go test -v ./keypairs/...
	@echo "Tests complete!"

test-xrpl:
	@echo "Running Go tests for xrpl package..."
	@go test -v ./xrpl/...
	@echo "Tests complete!"

test-ci:
	@echo "Running Go tests..."
	@go clean -testcache
	@go test -v $(EXCLUDED_TEST_PACKAGES) -parallel $(PARALLEL_TESTS) -timeout $(TEST_TIMEOUT)
	@echo "Tests complete!"

coverage-unit:
	@echo "Generating unit test coverage report..."
	@go test -coverprofile=coverage.out $(EXCLUDED_COVERAGE_PACKAGES)
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

benchmark:
	@echo "Running Go benchmarks..."
	@go test -bench=. ./...
	@echo "Benchmarks complete!"
