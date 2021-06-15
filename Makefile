.PHONY: build test
.DEFAULT_GOAL := build

build:
	go build -v ./cmd/app

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out