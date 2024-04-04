run:	
	go run main.go

init-app:
	go mod init authwithtoken
	go mod tidy

docker-restart:
	docker compose down -v
	docker compose up -d

migration-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:6432/postgres?sslmode=disable" -verbose up