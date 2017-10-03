include utils/help.mk

SOURCE_REPO?=gitlab.com/EasyStack/yakety

YAK?=bin/yak

bin:
	mkdir -p $@

.PHONY: build

build: ARGS?=-x
build: $(YAK)  ##@build build binary

$(YAK):
	go build -o $(YAK) $(ARGS) $(SOURCE_REPO)/cli/yak

.PHONY: test

test: ARGS?=-v
test: ##@source test
	go test $(ARGS) ./...

fmt: ARGS?=$(SOURCE_REPO)/...
fmt: ##@source fmt
	go fmt $(ARGS)

.PHONY: clean
clean:  ##@build remote build result
	rm -rf bin
