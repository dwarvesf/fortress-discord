dev:
	go run ./cmd/server/main.go

test:
	go test ./... --cover
