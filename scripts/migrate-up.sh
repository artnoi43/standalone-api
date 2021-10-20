#!/bin/sh
migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" up