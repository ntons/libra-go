TARGETS = echo_server image
ACTIONS = push

.PHONY: all $(TARGETS) $(ACTIONS)

all: $(TARGETS)

VERSION := $(shell cat VERSION)

echo_server:
	go build -o $@ github.com/ntons/libra-go/examples/echo_server

image:
	docker build -t ntons/echo_server:latest .
	docker tag ntons/echo_server:latest ntons/echo_server:$(VERSION)

push:
	docker push ntons/echo_server:latest
	docker push ntons/echo_server:$(VERSION)

