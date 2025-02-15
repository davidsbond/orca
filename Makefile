format:
	go tool buf format -w

generate:
	go generate ./...

migrate:
	go tool migrate create -dir internal/daemon/controller/database/migrations -ext sql $(NAME)
