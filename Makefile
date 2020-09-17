SHELL := /bin/bash

SRC = $(shell find . -type f -name '*.go' -not -iname '*.pb.*' -not -iname 'mock_*_test.go')

lint:
	@golangci-lint run ./...

test:
	@go test -cover ./...

cover:
	@go test -coverpkg=./internal/... -coverprofile=cover.out ./internal/... > /dev/null
	@go tool cover -func cover.out

build:
	@echo "  >  Building binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o scaffold cmd/scaffold/main.go

clean:
	@go clean

fmt:
	@gofmt -s -l -w $(SRC)

goimports:
	@goimports -w -local github.com/golangci/golangci-lint $(SRC)

format: fmt goimports
