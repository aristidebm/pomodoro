.PHONY: format lint install test

format:
	@go fmt .

lint:
	@echo "Linting ..."

install:
	@go mod tidy

test:
	@go test .
