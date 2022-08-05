gen: clean
	@echo "Generating mocks and API docs"
	go generate ./...
	swag init -g ./cmd/api/main.go -o ./docs/swagger
	swag fmt

clean:
	find . -type f \( -name '*_mock.go' -o -name '*_mock_test.go' \) -exec rm {} +

test:
	go test ./...

build:
	go build ./...

setup:
	docker-compose up -d

dah:
	git push heroku main

lah:
	heroku logs --tail

