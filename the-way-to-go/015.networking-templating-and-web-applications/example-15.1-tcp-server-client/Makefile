GO=go
GO_ENV=GOPATH=$(CURDIR)

PACKAGES=$(patsubst $(CURDIR)/src/%,%,$(shell find $(CURDIR)/src/ -maxdepth 1 -type d))
BINARY_SOURCES=$(wildcard $(CURDIR)/src/main_*.go)

install: test
	$(if $(wildcard $(CURDIR)/src/main.go), $(GO_ENV) $(GO) build -o bin/main src/main.go)
	@for SOURCE in $(BINARY_SOURCES) ; 		\
	do										\
		$(GO_ENV) $(GO) build -o $(CURDIR)/bin/$$( echo $$(basename -s .go $$SOURCE) | sed -e 's;^main_;;' ) $$SOURCE ; \
	done
	$(if $(PACKAGES), $(GO_ENV) $(GO) install$(foreach PACKAGE, $(PACKAGES), $(PACKAGE)))

test:
	$(if $(PACKAGES), $(GO_ENV) $(GO) test$(foreach PACKAGE, $(PACKAGES), $(PACKAGE)) -bench=.)

clean:
	rm -rf $(CURDIR)/bin/*
	find $(CURDIR) -name *.a -exec rm -rf "{}" \;
	find $(CURDIR) -name "log*" -exec rm -rf "{}" \;