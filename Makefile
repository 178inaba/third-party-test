.PHONY: all modtidy modernize lint test run

all: modtidy modernize lint test

modtidy:
	go mod tidy

modernize:
	go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -fix -test ./...

lint:
	docker compose run --rm lint --fix

test:
	go test -race -count 1 -cover ./...

run:
	go run main.go
