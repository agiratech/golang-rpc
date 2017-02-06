# golang-rpc

## Clone the repo for testing this example
`$ git clone git@github.com:agiratech/golang-rpc.git`

## Install gRPC
`$ go get google.golang.org/grpc`

## Install Protocol Buffers v3
`$ curl -OL https://github.com/google/protobuf/releases/download/v3.0.0-beta-2/protoc-3.0.0-beta-2-linux-x86_64.zip`

`$ unzip protoc-3.0.0-beta-2-linux-x86_64.zip -d protoc3`

`$ sudo mv protoc3/protoc /bin/protoc`

## Install the protoc plugin for Go
`$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
