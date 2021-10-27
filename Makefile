build:
	go build -v ./cmd/api

run:
	go run ./cmd/api/main.go -config-path=.env

test:
	go test -v -race ./...