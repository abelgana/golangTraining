.PHONY: all test build

all: update build test

update:
	bazel run //:gazelle update

test:
	bazel test --test_output=errors //...

build:
	bazel build //...
