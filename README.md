# Application Documentation

## Requirement
1. Go version 1.20 above
2. Docker for unit test
3. Swag, for api documentation. Installation command : go install github.com/swaggo/swag/cmd/swag@latest
4. Migrate, for database migration. https://github.com/golang-migrate/migrate
5. PostgresSQL Database
6. Mockery for mock in unit test, https://github.com/vektra/mockery

## Setup
1. Git clone this repository
2. Install application with command : 
```bash 
make init-app 
```
3. Copy appplication config with command : 
```bash 
cp config-example.yml config.yml
```
4. Edit `config.yml` file based on your own configuration.
5. Execute migration with command : 
```bash 
migrate -path migrations -database "postgresql://username:password@host:port/databasename?sslmode=disable" -verbose up
```

## Run Application
1. Run with command : 
```bash 
make run
```
2. API url example : [POST] localhost:7100/auth/register 
3. Open browser for api documentation : localhost:7100/swagger/

## Run Unit Test
1. Run docker with command : 
```bash 
make docker-restart
```
2. Run dummy data for testing with command :
```bash 
make test-migration-up
```
3. Run mock with command : 
```bash 
make mock-gen
```
4. Run unit test with command : 
```bash 
make test
```
5. To view coverage report in browser : 
```bash 
make test-coverage
```