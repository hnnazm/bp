include .env
export $(shell sed 's/=.*//' .env)

run:
	@go run *.go

build:
	@go build -o bin/server

migrateinstall:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrateup:
	@migrate -database ${DATABASE_DSN} -path ./migrations up

migratedown:
	@migrate -database ${DATABASE_DSN} -path ./migrations down

migratedrop:
	@migrate -database ${DATABASE_DSN} -path ./migrations drop -f
