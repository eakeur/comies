gen: clean
	go generate ./...

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

