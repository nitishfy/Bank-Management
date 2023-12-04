build:
	@go build -o ./bin/go

run: build
	@./bin/go

test:
	@go test -v ./...