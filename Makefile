.PHONY: all api

all:

API_PROTOS = $(wildcard ../libra-api/acct/v1/*.proto) \
			 $(wildcard ../libra-api/dbapi/v1/*.proto) \
			 $(wildcard ../libra-api/gwapi/v1/*.proto) \
			 $(wildcard ../libra-api/ptapi/v1/*.proto) \
			 $(wildcard ../libra-api/sdk/v1/*.proto)

api:
	mkdir -p api && rm -rf api/*
	for PROTO in $(API_PROTOS); do \
		protoc -I../libra-api \
		--go_out=api \
		--go_opt=paths=source_relative \
		--go-grpc_out=api \
		--go-grpc_opt=paths=source_relative \
		$$PROTO; done

