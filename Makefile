compile-proto:
	protoc --go_out=. --dart_out=grpc:views/comies/lib/gateway/rpc --proto_path ./protos protos/*.proto google/protobuf/timestamp.proto