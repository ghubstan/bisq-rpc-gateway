generate:

    # Generate client, server & reverse-proxy stubs
	protoc \
		-I /usr/local/include -I . \
		-I $(GOPATH)/src \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc,paths=source_relative:. \
		--grpc-gateway_out=logtostderr=true:. \
		proto/service.proto


install:

	GO111MODULE=on go get \
        github.com/golang/protobuf/protoc-gen-go \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

build:

	go build ./...

all:  install generate build