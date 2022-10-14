.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golint ./...
.PHONY:lint

vet: fmt
		go vet ./...
.PHONY:vet

build: build
		go build cmd/main.go
.PHONY:build	