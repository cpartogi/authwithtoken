run:	
	go run main.go

docker-restart:
	docker compose down -v
	docker compose up -d

migration-up:
	migrate -path migration -database "postgresql://postgres:postgres@localhost:6432/postgres?sslmode=disable" -verbose up