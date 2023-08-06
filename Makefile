linux-bin:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/go-todo-api ./cmd/go-todo-api/main.go

run-watch:
	$(source .envrc)
	go generate ./...
	air -c ./.air.toml

test:
	./scripts/run-test.sh