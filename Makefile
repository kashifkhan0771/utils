# Run tests
test:
	bash ./scripts/test.sh

# Run golangci-lint
lint:
	@command -v golangci-lint > /dev/null || (echo "Error: golangci-lint is not installed" && exit 1)
	golangci-lint run
