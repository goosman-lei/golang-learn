GO=go
GO_ENV=GOPATH=$(CURDIR)

all: pkg-udsort main

main:
	$(GO_ENV) $(GO) build -o bin/main src/main.go

pkg-udsort:
	$(GO_ENV) $(GO) install udsort