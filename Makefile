# settings
IS_RELEASE ?= false

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# tools
GO = go
WINRES = $(GO) run github.com/tc-hib/go-winres@latest
UPX = upx

# output
BINARY = ./bin/$(GOOS)-$(GOARCH)

EXE_EXT = $(shell go env GOEXE)
ifeq ($(GOARCH),wasm)
	EXE_EXT = .wasm
endif
EXE = $(BINARY)/gophengine$(EXE_EXT)

# flags
UPX_FLAGS = --best --lzma

GE_FLAGS ?=

GO_GCFLAGS =
GO_LDFLAGS =
GO_FLAGS = -v

ifeq ($(IS_RELEASE),true)
	GO_GCFLAGS += -dwarf=false
	GO_LDFLAGS += -s -w
	GO_FLAGS += -trimpath
	ifeq ($(GOOS),windows)
	GO_LDFLAGS += -H windowsgui
	endif
endif

GO_FLAGS += -gcflags="$(GO_GCFLAGS)" -ldflags="$(GO_LDFLAGS)" -buildvcs=true

ifneq ($(GOARCH),wasm)
	GO_FLAGS += -buildmode=pie
endif

.PHONY: all
all: $(EXE) upx

.PHONY: run
run:
	$(GO) get
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) ./cmd/gophengine $(GE_FLAGS)

.PHONY: run-debug
run-debug:
	$(GO) get
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) ./cmd/gophengine --log-level=debug $(GE_FLAGS)

$(EXE): clean
	mkdir -p $(BINARY)

	$(GO) get
ifeq ($(GOOS),windows)
	$(WINRES) make  --out ./cmd/gophengine/rsrc
endif
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(EXE) ./cmd/gophengine

.PHONY: upx
upx: $(EXE)
ifeq ($(IS_RELEASE),true)
	ifneq ($(GOARCH),wasm)
		$(UPX) $(UPX_FLAGS) $(EXE)
	endif
endif

.PHONY: clean
clean:
	rm -rf $(BINARY)
ifeq ($(GOOS),windows)
	rm -f ./cmd/gophengine/rsrc_windows_*.syso
endif
