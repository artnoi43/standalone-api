#!/bin/sh
docker run -p 5432:5432 --name postgres13 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:13-alpine