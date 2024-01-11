# settings
IS_RELEASE = false

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# tools
GO = go
WINRES = $(GO) run github.com/tc-hib/go-winres@latest

# icon
ICON = icon.ico

# output
BINARY = ./bin/$(GOOS)-$(GOARCH)

EXE = $(BINARY)/gophengine.exe

# flags
UPX_FLAGS = --best --lzma

GO_GCFLAGS =
GO_LDFLAGS =
GO_FLAGS = -v

ifeq ($(IS_RELEASE),true)
	GO_GCFLAGS += -dwarf=false
	GO_LDFLAGS += -s -w
endif

GO_FLAGS += -gcflags="$(GO_GCFLAGS)" -ldflags="$(GO_LDFLAGS)" -buildvcs=true

.PHONY: all
all: build upx

.PHONY: build
build: clean
	mkdir -p $(BINARY)

	$(GO) get
	$(WINRES) make
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(BINARY)

.PHONY: upx
upx:
ifeq ($(IS_RELEASE),true)
	upx $(UPX_FLAGS) $(EXE)
endif

.PHONY: clean
clean:
	rm -rf $(BINARY)/$(GOOS)-$(GOARCH)
	rm -f rsrc_windows_386.syso
	rm -f rsrc_windows_amd64.syso