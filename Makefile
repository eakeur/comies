compile-proto:
	protoc --go_out=. --dart_out=grpc:views/comies/lib/gateway/rpc --proto_path ./protos protos/*.proto google/protobuf/timestamp.proto

generate: clean-gen
	go generate ./...

clean-gen:
	find . -type f \( -name '*_mock.go' -o -name '*_mock_test.go' \) -exec rm {} +
	rm -rf gen/*

test:
	go test ./...