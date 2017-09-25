SOURCE_REPO?=gitlab.com/EasyStack/yakety

YAK?=bin/yak

bin:
	mkdir -p $@

.PHONY: build

build: ARGS?=-x
build: $(YAK)

$(YAK):
	go build -o $(YAK) $(ARGS) $(SOURCE_REPO)/cli/yak

.PHONY: clean
clean:
	rm -rf bin
