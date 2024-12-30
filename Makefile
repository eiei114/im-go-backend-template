.PHONY: format tidy install-tools lint all

format:
	cd app && go fmt ./...

tidy:
	cd app && go mod tidy

test:
	cd app && go test ./...

all: format tidy test

exec:
	cd app && docker compose -f compose.yml up -d && go run cmd/main.go

down:
	cd app && docker compose -f compose.yml down
