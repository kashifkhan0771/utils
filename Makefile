# Run tests
test:
	bash ./scripts/test.sh

lint:
	@which golangci-lint > /dev/null || (echo "Error: golangci-lint is not installed" && exit 1)
	golangci-lint run --config=.golangci.yml
