.PHONY: test test-race coverage

# Run all tests
test:
	go clean -testcache && go test ./... -v -failfast

# Run tests with race detection
test-race:
	go clean -testcache && go test ./... -race -v -failfast

# Run tests with coverage
coverage:
	go test ./... -coverprofile=coverage.out -v -failfast
	go tool cover -func=coverage.out

# Run tests for a specific package
test-package:
	go clean -testcache && go test $(PACKAGE) -v -failfast

# Clean test cache
test-clean:
	go clean -testcache