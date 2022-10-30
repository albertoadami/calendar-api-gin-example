.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint:
		golint ./...
.PHONY:lint

vet:
		go vet ./...
.PHONY:vet

test:
		go test ./...
.PHONY:test

build: build
		go build cmd/main.go
.PHONY:build	