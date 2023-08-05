GO=go
RSRC=rsrc

ICON=icon.ico
BINARY=./bin
SYSO=./GophEngine$(shell $(GO) env GOEXE).syso

FLAGS=-v

GOOS=windows
GOARCH=amd64

all: build

build: clean
	mkdir -p $(BINARY)/$(GOOS)-$(GOARCH)

	$(GO) get

	$(RSRC) -ico="$(ICON)" -o $(SYSO)
	CGO_ENABLED=1 GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(FLAGS) -o $(BINARY)/$(GOOS)-$(GOARCH)

clean:
	rm -rf $(BINARY)/$(GOOS)-$(GOARCH)
	rm $(SYSO)