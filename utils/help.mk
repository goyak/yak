# Refer HELP_FUN
# https://gist.github.com/prwhite/8168133
# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

PERL :=  $(shell command -v perl 2> /dev/null)
HELP_FUN = \
    %help; \
    %help_info; \
    while(<>) { \
        if(/^([a-z0-9_-]+):.*\#\#(?:@([\w-]+))?\s(.*)$$/) { \
            push(@{$$help{$$2}}, [$$1, $$3]); \
        } \
        if(/^\#\#(?:@(\w+))?\s(.*)$$/) { \
            push(@{$$help_info{$$1}}, $$2); \
        } \
    }; \
    print "usage: make [target]\n\n"; \
    for ( sort keys %help ) { \
        printf("%-21s \033[33m %5s \033[0m\n", $$_.":", @{$$help_info{$$_}});\
        printf(" \033[36m %-20s \033[0m %s\n", $$_->[0], $$_->[1]) for @{$$help{$$_}}; \
        print "\n"; \
    }

help:
ifdef PERL
	@$(PERL) -e '$(HELP_FUN)' $(MAKEFILE_LIST)
else
	@grep -E -h '^[a-zA-Z0-9_-]+:.*?##@.*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?##@[a-zA-Z0-9_-]* "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
endif
