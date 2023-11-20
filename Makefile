build:
	@go build -C cmd -o ../bin/accord

run: clean build
	@./bin/accord

test:
	@go test

clean:
	@go clean

.PHONY: build run test clean