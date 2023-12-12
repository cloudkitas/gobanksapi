build:
	@go build -o bin/gobanks

run: build
	@./bin/gobanks

test:
	@go test -v ./...
