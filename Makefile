# Run tests
test:
	bash ./scripts/test.sh

# Run golangci-lint
lint:
	@command -v golangci-lint > /dev/null || (echo "Error: golangci-lint is not installed" && exit 1)
	@version=$$(golangci-lint version | grep -oE 'version [0-9]+\.[0-9]+\.[0-9]+' | awk '{print $$2}'); \
	config=".golangci.yml"; \
	if echo "$$version" | grep -q '^1\.'; then \
    		config=".golangci.v1.yml"; \
	fi; \
	golangci-lint run --config=$$config
