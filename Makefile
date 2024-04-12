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

mock-gen:
	mockery --dir domain/auth --name AuthUsecaseInterface --filename iauth_usecase.go --output domain/auth/mocks --with-expecter
	mockery --dir domain/auth --name AuthRepoInterface --filename iauth_repo.go --output domain/auth/mocks --with-expecter	

test:
	go test -p 1 --v authwithtoken/domain/auth/testcase -coverprofile cover.out -coverpkg authwithtoken/domain/auth/usecase
	go tool cover -func cover.out	

test-coverage:
	go tool cover -html cover.out		