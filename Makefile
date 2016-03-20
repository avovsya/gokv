GO ?= go

.PHONY: all
all: build start

start:
	./gokv

build:
	$(GO) build
