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

mod:
	go mod download
.PHONY:mod

build: vet mod
	go build
.PHONY:build

clean:
	go clean
.PHONY:clean