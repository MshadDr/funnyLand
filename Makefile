serv:
	go run main.go serv

migrate-up:
	go run ./cmd/migrate/main.go -up

migrate-down:
	go run ./cmd/migrate/main.go -down

.PHONY: serv migrate-up migrate-down
