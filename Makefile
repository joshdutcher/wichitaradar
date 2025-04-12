.PHONY: test

# Run all tests
test:
	go test ./... -v

# Run tests with race detection
test-race:
	go test ./... -race -v

# Clean test cache
test-clean:
	go clean -testcache