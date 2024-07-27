generate:
	go run github.com/99designs/gqlgen generate

server:
	go run server.go

.PHONY: generate
