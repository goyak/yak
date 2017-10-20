.DEFAULT_GOAL := help

SOURCE_REPO ?= gitlab.com/EasyStack/yakety
YAK ?= bin/yak
YAKD ?= bin/yakd


include utils/help.mk

bin:
	@mkdir -p $@

.PHONY: build

build: ARGS?=-x
build: ##@build build binary
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

.PHONY: test

test: ARGS?=+local
test: ##@source to run specific unittests $ ARGS=gitlab.com/EasyStack/yakety/lib/index make test
	@govendor test -v $(ARGS)

fmt: ARGS?=$(SOURCE_REPO)/...
fmt: ##@source fmt
	@govendor fmt +local
