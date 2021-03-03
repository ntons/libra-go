TARGETS = greeter_server image
ACTIONS = push

.PHONY: all $(TARGETS) $(ACTIONS)

all: $(TARGETS)

VERSION := $(shell cat VERSION)

greeter_server:
	go build -o $@ github.com/ntons/libra-go/examples/greeter_server

image:
	docker build -t ntons/greeter_server:latest .
	docker tag ntons/greeter_server:latest ntons/greeter_server:$(VERSION)

push:
	docker push ntons/greeter_server:latest
	docker push ntons/greeter_server:$(VERSION)

