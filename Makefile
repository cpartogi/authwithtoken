run:
	swag init	
	go run main.go

init-app:
	go mod init authwithtoken
	go mod tidy

docker-restart:
	docker compose down -v
	docker compose up -d

test-migration-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:6432/postgres?sslmode=disable" -verbose up

test-migration-down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:6432/postgres?sslmode=disable" -verbose down

mock-gen:
	mockery --dir domain/auth --name AuthRepoInterface --filename iauth_repo.go --output domain/auth/mocks --with-expecter	

test:
	go test -p 1 --v authwithtoken/domain/auth/testcase -coverprofile cover.out -coverpkg authwithtoken/domain/auth/usecase,authwithtoken/domain/auth/repo
	go tool cover -func cover.out	

test-coverage:
	go tool cover -html cover.out		