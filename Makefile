.PHONY: all api

all:

api: $(shell find ../libra-api -name "*.proto")
	mkdir -p api/libra && rm -rf api/libra/*
	echo "Date:   `date +'%Y-%m-%d %H:%M:%S'`" >> api/libra/VERSION
	echo "Commit: `git --git-dir=../libra-api/.git --work-tree=../libra-api/ rev-list -1 HEAD`" >> api/libra/VERSION
	for PROTO in $^; do \
		protoc -I../libra-api \
		--go_out=api \
		--go_opt=paths=source_relative \
		--go-grpc_out=api \
		--go-grpc_opt=paths=source_relative \
		$$PROTO; done
	for PROTO in $$(grep -H google.api.http $^|cut -d: -f1|sort|uniq); do \
		protoc -I../libra-api \
		--grpc-gateway_out=paths=source_relative:api \
		$$PROTO; done

