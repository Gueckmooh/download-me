-include config.mk
.SECONDEXPANSION:

BUILDDIR ?= $(CURDIR)
RESOURCES ?= $(BUILDDIR)/resources

BINDIR      ?= $(BUILDDIR)/bin
JSDIR       ?= $(RESOURCES)/public/js
ALLBIN ?= $(notdir $(shell find cmd -mindepth 1 -type d -print))

# go option
PKG        := ./...
TAGS       :=
TESTS      := .
TESTFLAGS  :=
LDFLAGS    := -w -s
GOFLAGS    :=

DEPDIR ?= .deps
.PRECIOUS: %/.f $(BUILDDIR)/$(DEPDIR)/%.d

%/.f:
	$(QUIET)mkdir -p $(dir $@)
	$(QUIET)touch $@

NOINC = clean, mrproper

SRC := $(shell find . -type f -name '*.go' -print) go.mod
DARTSRC := $(shell find . -type f -name '*.dart' -exec grep -l 'void main' {} +)

ALLBINS := $(addprefix $(BINDIR)/, $(ALLBIN))
ALLJS := $(addprefix $(JSDIR)/, $(addsuffix .js, $(notdir $(basename $(DARTSRC)))))

.DEFAULT_GOAL := all

-include $(addprefix $(DEPDIR)/, $(patsubst %.dart,\
                                     %.d,$(notdir $(DARTSRC))))

.PHONY: all
all: build build-js

# Required for globs to work correctly
SHELL      = /usr/bin/env bash

# ------------------------------------------------------------------------------
#  build

.PHONY: build
build: $(ALLBINS)

$(BINDIR)/%: $(SRC)
	GO111MODULE=on go build $(GOFLAGS) -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BINDIR)'/$(BINNAME) ./cmd/$(notdir $@)

.PHONY: build-js
build-js: $(ALLJS)

$(BUILDDIR)/$(DEPDIR)/%.d: pkg/client/%.dart $$(@D)/.f
	./scripts/dartdeps.sh $< $(addprefix $(JSDIR)/, $(notdir $(basename $<).js)) > $@

$(JSDIR)/%.js: pkg/client/%.dart $(BUILDDIR)/$(DEPDIR)/%.d $$(@D)/.f
	dart2js $< -o $@

