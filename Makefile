.PHONY: test test-race coverage lint lint-fix

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

# Run linters (Go + JavaScript)
lint:
	@echo "Running Go linter..."
	golangci-lint run ./...
	@echo "Running JavaScript linter..."
	npm run lint

# Run linters with auto-fix
lint-fix:
	@echo "Running Go linter with auto-fix..."
	golangci-lint run ./... --fix
	@echo "Running JavaScript linter with auto-fix..."
	npm run lint:fix