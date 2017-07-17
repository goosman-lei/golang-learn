GO=go

install: package binary

clean:
	$(GO) clean

binary: bin-ucmain

package: pkg-uc

bin-ucmain:
	$(GO) build -o bin/ucmain src/ucmain.go

pkg-uc:
	$(GO) install uc