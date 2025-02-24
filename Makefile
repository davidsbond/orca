format:
	go tool buf format -w
	go fmt ./...

generate:
	go generate ./...

migrate:
	go tool migrate create -dir internal/daemon/controller/database/migrations -ext sql $(NAME)

test:
	go test -race -short ./...

lint:
	go vet ./...

release:
	go tool goreleaser release --clean

snapshot:
	go tool goreleaser release --snapshot
