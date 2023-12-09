all: protoc.pb protoc.gw

protoc.pb:
	protoc \
	-I ./proto \
	--go_out=./pb/ \
	--go-grpc_out=./pb/ \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./proto/greeter.proto

protoc.gw:
	protoc \
	-I ./proto \
	--grpc-gateway_out ./pb \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./proto/greeter.proto ./proto/google/api/httpbody.proto ./proto/google/api/http.proto ./proto/google/api/field_behavior.proto ./proto/google/api/annotations.proto

