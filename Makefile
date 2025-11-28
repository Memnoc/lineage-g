.PHONY: build test run clean install

BINARY_NAME=lineage

build:
	go build -o bin/$(BINARY_NAME) ./cmd/lineage

test:
	go test -v ./...

run:
	go run ./cmd/lineage

clean:
	rm -rf bin/
	rm -f recipes.typ recipes.pdf

install:
	go install ./cmd/lineage

fmt:
	go fmt ./...

lint:
	golangci-lint run

dev: fmt test build

.DEFAULT_GOAL := build
