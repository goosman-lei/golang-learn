GO=go
GO_ENV=GOPATH=$(CURDIR)

PACKAGES=even

install: test
	$(GO_ENV) $(GO) build -o bin/main src/main.go
	$(GO_ENV) $(GO) install$(foreach PACKAGE, $(PACKAGES), $(PACKAGE))

test:
	$(GO_ENV) $(GO) test$(foreach PACKAGE, $(PACKAGES), $(PACKAGE)) -bench=.

clean:
	rm -rf $(CURDIR)/bin/*
	find $(CURDIR) -name *.a -exec rm -rf "{}" \;
	find $(CURDIR) -name "log*" -exec rm -rf "{}" \;