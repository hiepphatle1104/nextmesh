.PHONY: build run server client docker docker-up docker-down

build:
	@go build -o .bin/server cmd/main.go

server: build
	@.bin/server

client:
	@cd client && npm run dev

run: server client

docker-up:
	@docker compose -f docker-compose.yml up -d

docker-down:
	@docker compose -f docker-compose.yml down
