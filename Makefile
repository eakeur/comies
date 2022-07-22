gen: clean
	go generate ./...

clean:
	find . -type f \( -name '*_mock.go' -o -name '*_mock_test.go' \) -exec rm {} +

test:
	go test ./...

setup:
	docker-compose up -d

deploy-api-homolog:
	gcloud app deploy hml.yaml

log-api-homolog:
	gcloud app logs tail -s api-homolog

