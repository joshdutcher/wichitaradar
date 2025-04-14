.PHONY: test test-race test-clean test-coverage test-package

# Run all tests
test:
	go test ./... -v

# Run tests with race detection
test-race:
	go test ./... -race -v

# Run tests with coverage
test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

# Run tests for a specific package
test-package:
	go test $(PACKAGE) -v

# Clean test cache
test-clean:
	go clean -testcache