build:
	@go build -C server/cmd -o ../../bin/accord

run: build
	@./bin/accord

test:
	@go test

clean:
	@go clean


db-up:
	@migrate -path server/db/migrations -database "postgres://root:test@localhost:5432/accord-chat?sslmode=disable" -verbose up

db-down:
	@migrate -path server/db/migrations -database "postgres://root:test@localhost:5432/accord-chat?sslmode=disable" -verbose down

.PHONY: build run test clean db-up db-down