build:
	go build -o ./pbserver ./cmd/server.go
	go build -o ./pbclient ./cmd/client.go