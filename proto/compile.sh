#!/usr/bin/env bash -e

# Install protobuf
#  brew install protobuf
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://grpc.io/

cd $(dirname $0)

for proto in *.proto
do
  echo "generating Go for $proto"
  protoc \
    -I . \
    --go_out=../pb/ \
    --go-grpc_out=../pb/ \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    $proto
done
