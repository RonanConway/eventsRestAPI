build:
	@go build -o bin/eventsRestApi

run: build
	@./bin/eventsRestApi

test:
	go test -v ./...