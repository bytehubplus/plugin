.PHONY: all

all: 
	protoc --go_out=protos/core \
	--go_opt=module=github.com/bytehubplus/plugin/protos/core \
	--go-grpc_out=protos/core \
	--go-grpc_opt=module=github.com/bytehubplus/plugin/protos/core \
	--go-grpc_opt=require_unimplemented_servers=false \
	./protos/core/base.proto

