GO=go
WINRES=go-winres

ICON=icon.ico
BINARY=./bin

FLAGS=-v

GOOS=windows
GOARCH=amd64

all: build

build: clean
	mkdir -p $(BINARY)/$(GOOS)-$(GOARCH)

	$(GO) get
	$(WINRES) make
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(FLAGS) -o $(BINARY)/$(GOOS)-$(GOARCH)

clean:
	rm -rf $(BINARY)/$(GOOS)-$(GOARCH)
	rm rsrc_windows_386.syso
	rm rsrc_windows_amd64.syso