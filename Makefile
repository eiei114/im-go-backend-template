.PHONY: format tidy install-tools lint all

format:
	cd app && go fmt ./...

tidy:
	cd app && go mod tidy

test:
	cd app && go test ./...

all: format tidy test