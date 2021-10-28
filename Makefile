build:
	go build -v ./cmd/api

run:
	go run ./cmd/api/main.go -config-path=.env

test:
	go test -v -race ./...

db-run:
	docker run --name mysql -e MYSQL_DATABASE=db -e MYSQL_USER=user -e MYSQL_PASSWORD=secret -e MYSQL_ROOT_PASSWORD=secret -p 3306:3306 -d --rm mysql
db-kill:
	docker kill mysql