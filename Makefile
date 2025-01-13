# Run tests
test:
	bash ./scripts/test.sh

lint:
	golangci-lint run
