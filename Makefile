GO=go
WINRES=go-winres

ICON=icon.ico
BINARY=./bin

FLAGS=-v

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p $(BINARY)/$(GOOS)-$(GOARCH)

	$(GO) get
	$(WINRES) make
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(FLAGS) -o $(BINARY)/$(GOOS)-$(GOARCH)

.PHONY: clean
clean:
	rm -rf $(BINARY)/$(GOOS)-$(GOARCH)
	rm -f rsrc_windows_386.syso
	rm -f rsrc_windows_amd64.syso