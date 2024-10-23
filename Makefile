# You can check here for a most complete Makefile example
# https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects

.PHONY: format lint install test run

format:
	@go fmt ./...

lint:
	@echo "Linting ..."

install:
	@go mod tidy

test:
	@go test ./...

run:
	@go run main.go
