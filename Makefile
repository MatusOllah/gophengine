# settings
IS_RELEASE ?= false

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# tools
GO = go
WINRES = $(GO) run github.com/tc-hib/go-winres@latest
UPX = upx
WASM_OPT = wasm-opt

# output
BINARY = ./bin/$(GOOS)-$(GOARCH)

EXE_EXT = $(shell go env GOEXE)
ifeq ($(GOARCH),wasm)
	EXE_EXT = .wasm
endif
EXE = $(BINARY)/gophengine$(EXE_EXT)

WASM_OPT_OUT = $(EXE:.wasm=.opt.wasm)

# flags
GO_GCFLAGS =
GO_LDFLAGS =
GO_FLAGS = -v

UPX_FLAGS = -f --best --lzma
WASM_OPT_FLAGS = -Oz --strip-debug --strip-producers --enable-bulk-memory-opt

GE_FLAGS ?=

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
all: clean build

.PHONY: run
run:
	$(GO) get
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) ./cmd/gophengine $(GE_FLAGS)

.PHONY: run-debug
run-debug:
	$(GO) get
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) ./cmd/gophengine --log-level=debug $(GE_FLAGS)

.PHONY: build
build: $(BINARY)

$(BINARY):
	mkdir -p $(BINARY)

	$(GO) get
ifeq ($(GOOS),windows)
	$(WINRES) make --out ./cmd/gophengine/rsrc
endif
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(EXE) ./cmd/gophengine

ifeq ($(IS_RELEASE),true)
ifneq ($(GOARCH),wasm)
	strip $(EXE)
	$(UPX) $(UPX_FLAGS) $(EXE)
endif
ifeq ($(GOARCH),wasm)
	$(WASM_OPT) $(WASM_OPT_FLAGS) -o $(WASM_OPT_OUT) $(EXE)
	rm $(EXE)
	mv $(WASM_OPT_OUT) $(EXE)
endif
endif

.PHONY: clean
clean:
	rm -rf $(BINARY)
ifeq ($(GOOS),windows)
	rm -f ./cmd/gophengine/rsrc_windows_*.syso
endif

.PHONY: clean-all
clean-all:
	rm -rf ./bin/
ifeq ($(GOOS),windows)
	rm -f ./cmd/gophengine/rsrc_windows_*.syso
endif
