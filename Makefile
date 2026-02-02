.PHONY: run test docker-up docker-down

run:
	go run ./cmd/api

test:
	go test ./...

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down -v
