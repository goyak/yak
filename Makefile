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

.PHONY: clean
clean:  ##@build remote build result
	rm -rf bin
