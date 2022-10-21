test:
	go test ./...

up:
	docker-compose up -d

serve: up
	go run cmd/api/main.go