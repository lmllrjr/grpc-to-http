#!/usr/bin/env bash -e

# https://github.com/grpc-ecosystem/grpc-gateway
#
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
		--grpc-gateway_out ../pb \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    $proto ./google/api/httpbody.proto ./google/api/http.proto ./google/api/field_behavior.proto ./google/api/annotations.proto
done
