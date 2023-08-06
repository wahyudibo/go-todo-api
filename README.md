# go-todo-api

My sample app to demonstrates (also for documentation purpose) about best practices that i learned along the way when developing RESTful API using golang.

## Prerequisites

- Go v1.19
- MySQL v8
- Make
- [air](https://github.com/cosmtrek/air). When encountered with `GLIBC_2.34 not found` error, make sure to install GLIBC_2.34 or latest.
- AWS account for s3 or any AWS services. Initializes using `aws configure --profile=YOUR_PROFILE_NAME` command and make sure it has setup by running `printenv | grep AWS_PROFILE` and it returns your profile name

## Starting the app

- Make sure that you already setup the database connection configuration based on `.envrc` file.
- Run the program using `make run-watch`