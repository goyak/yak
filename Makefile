.DEFAULT_GOAL := help

SOURCE_REPO ?= github.com/goyak/yak
YAK ?= bin/yak
YAKD ?= bin/yakd


include utils/help.mk

bin:
	@mkdir -p $@

.PHONY: build

build: ARGS?=-x
build: ##@build build binary, $ FORCE=1 make build
ifdef FORCE
	@make clean
endif
	@ARGS=$(ARGS) make $(YAK) $(YAKD)
$(YAK):
	@go build -o $(YAK) $(ARGS) $(SOURCE_REPO)/cli/yak
$(YAKD):
	@go build -o $(YAKD) $(ARGS) $(SOURCE_REPO)/cli/yakd

.PHONY: clean
clean: ##@build remote build result
	@rm -rf bin

install: ##@build install binary into GOPATH/bin
	@govendor install -v +local

install_clean: ##@build remove installed binary
	rm $(GOPATH)/bin/yak
	rm $(GOPATH)/bin/yakd

.PHONY: test

test: ARGS?=+local
test: ##@source to run specific unittests $ ARGS="github.com/goyak/yak/lib/index -v" make test
	@govendor test --cover -v $(ARGS)

fmt: ARGS?=$(SOURCE_REPO)/...
fmt: ##@source fmt
	@govendor fmt +local
