build:
	go build -v ./cmd/api

run:
	go run ./cmd/api/main.go

test:
	go test -v -race ./...