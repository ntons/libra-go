.PHONY: all api

all:

API_PROTOS =  \
			 $(shell find libra-api/db/         -name "*.proto") \
			 $(shell find libra-api/gw/         -name "*.proto") \
			 $(shell find libra-api/pt/         -name "*.proto") \
			 $(shell find libra-api/sdk/        -name "*.proto") \
			 $(shell find libra-api/thirdparty/ -name "*.proto")

api: $(API_PROTOS)
	mkdir -p api && rm -rf api/*
	for PROTO in $^; do \
		protoc -Ilibra-api \
		--go_out=api \
		--go_opt=paths=source_relative \
		--go-grpc_out=api \
		--go-grpc_opt=paths=source_relative \
		$$PROTO; done
	for PROTO in $$(grep -H google.api.http $^|cut -d: -f1|sort|uniq); do \
		protoc -Ilibra-api \
		--grpc-gateway_out=paths=source_relative:api \
		$$PROTO; done

