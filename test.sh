#!/bin/bash

# Pull the PostgreSQL Docker image
docker pull postgres

# Create and start the Docker container
docker run --name mock-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

# Run your tests here, assuming your test command is "go test"
go test ./...

# Stop and remove the Docker container
docker stop mock-postgres
docker rm mock-postgres