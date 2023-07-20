PB_ROOT = ./pb

proto_files_proto = ${notdir ${wildcard ${PB_ROOT}/proto/*.proto}}
proto_files = ${proto_files_proto:%.proto=%}

.PHONY: all
all: proto lint

proto: ${proto_files}

lint:
	golangci-lint run --config=./.golangci.yml

%: ${PB_ROOT}/proto/%.proto
	protoc -I ${PB_ROOT}/proto/ --go-grpc_out=${PB_ROOT}/gen --go_out=${PB_ROOT}/gen/ ${PB_ROOT}/proto/$@.proto

