.PHONY: all api

all:

api: $(shell find protos/libra-api ! -path '*/google/*' -name "*.proto")
	mkdir -p api && rm -rf api/*
	for PROTO in $^; do \
		protoc -Iprotos/libra-api \
		--go_out=api \
		--go_opt=paths=source_relative \
		--go-grpc_out=api \
		--go-grpc_opt=paths=source_relative \
		$$PROTO; done
	for PROTO in $$(grep -H google.api.http $^|cut -d: -f1|sort|uniq); do \
		protoc -Iprotos/libra-api \
		--grpc-gateway_out=paths=source_relative:api \
		$$PROTO; done

