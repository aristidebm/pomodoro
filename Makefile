.PHONY: format lint install test run

format:
	@go fmt .

lint:
	@echo "Linting ..."

install:
	@go mod tidy

test:
	@go test .

run:
	@go run main.go
