# go-todo-api
Go Todo API provides sample todo api using go-chi and mysql (via gorm)

## Prerequisites
- Go v1.19
- MySQL v8

## Starting the app
- Make sure that you already setup the database connection configuration based on `.envrc` file.
- Run `source .envrc` to export all environment variables to your shell.
- Run the program using `go run cmd/go-todo-api/main.go`