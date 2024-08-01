.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/server/

run: build
	./cmd/server/server

clean:
	go clean -i
