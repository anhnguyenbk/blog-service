FUNCTIONNAME="blog-post-function"

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

install:
	go mod download

build:
	GOOS=linux GOARCH=amd64 go build -o $(GOBIN)/main ./cmd/post/main.go || exit
	
build-local:
	go build -o $(GOBIN)/main ./cmd/post/main.go

deploy:
	zip -rj main.zip bin/main
	aws lambda update-function-code --function-name $(FUNCTIONNAME) --zip-file fileb://main.zip

deploy-local:
	./bin/main
