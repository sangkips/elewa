build:
	@go build -o bin/go-elewa

run: build
	@./bin/go-elewa

test:
	@go test -v ./...