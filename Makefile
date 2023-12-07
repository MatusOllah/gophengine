# settings
ISRELEASE = true

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# tools
GO = go
WINRES = $(GO) run github.com/tc-hib/go-winres@latest

# icon
ICON = icon.ico

# output
BINARY = ./bin

# flags
GO_GCFLAGS =
GO_LDFLAGS =
GO_FLAGS = -v

ifeq ($(ISRELEASE),true)
	GO_GCFLAGS += -dwarf=false
	GO_LDFLAGS += -s -w
endif

GO_FLAGS += -gcflags="$(GO_GCFLAGS)" -ldflags="$(GO_LDFLAGS)"

.PHONY: all
all: build

.PHONY: build
build: clean
	mkdir -p $(BINARY)/$(GOOS)-$(GOARCH)

	$(GO) get
	$(WINRES) make
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(BINARY)/$(GOOS)-$(GOARCH)

.PHONY: clean
clean:
	rm -rf $(BINARY)/$(GOOS)-$(GOARCH)
	rm -f rsrc_windows_386.syso
	rm -f rsrc_windows_amd64.syso