PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

install:
	go mod download

build:
	go build -o $(GOBIN)/main ./cmd/post/main.go || exit