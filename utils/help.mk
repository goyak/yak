# Refer HELP_FUN
# https://gist.github.com/prwhite/8168133

##@miscellaneous description for catalog

UID = $(shell id -u)
SUDO_UID ?= 0
SUDO_GID ?= 0

HELP_FUN = \
    %help; \
    %help_info; \
    while(<>) { \
        if(/^([a-z0-9_-]+):.*\#\#(?:@(\w+))?\s(.*)$$/) { \
            push(@{$$help{$$2}}, [$$1, $$3]); \
        } \
        if(/^\#\#(?:@(\w+))?\s(.*)$$/) { \
            push(@{$$help_info{$$1}}, $$2); \
        } \
    }; \
    print "usage: make [target]\n\n"; \
    for ( sort keys %help ) { \
        printf("%-21s \033[33m %5s \033[0m\n", $$_.":", @{$$help_info{$$_}}); \
        printf(" \033[36m %-20s \033[0m %s\n", $$_->[0], $$_->[1]) for @{$$help{$$_}}; \
        print "\n"; \
    }

help:	##@miscellaneous shows common make targets
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)
