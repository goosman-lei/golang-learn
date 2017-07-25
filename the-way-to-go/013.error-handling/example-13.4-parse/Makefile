GO=go
GO_ENV=GOPATH=$(CURDIR)

all: pkg-parse main

main:
	$(GO_ENV) $(GO) build -o bin/main src/main.go

pkg-parse:
	$(GO_ENV) $(GO) install parse