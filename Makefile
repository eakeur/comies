PROTO_REPO=git@github.com:eakeur-institute/comies-brunch.git

protos:
	rm -rf protos
	git clone ${PROTO_REPO} protos
	protoc --go_out=. --go-grpc_out=. --proto_path ./protos protos/*/*.proto
	rm -rf protos

gen: clean protos
	go generate ./...

clean:
	find . -type f \( -name '*_mock.go' -o -name '*_mock_test.go' \) -exec rm {} +
	find . -type f \( -name '*pb.go' -o -name '*pb_test.go' \) -exec rm {} +
	rm -rf protos
	rm -rf gen/*

test:
	go test ./...