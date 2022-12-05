gen: clean gogen docgen

gogen:
	go generate ./...

docgen:
	cd cmd/api && swag init -g ./main.go -o ../../docs/swagger
	cd cmd/api && swag fmt

clean:
	find . -type f \( -name '*_mock.go' -o -name '*_mock_test.go' \) -exec rm {} +

test: gen
	go test ./...

build:
	go build ./...

up:
	docker-compose up -d

