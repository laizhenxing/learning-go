#!/usr/bin/bash

go run main.go --server_address :8081 &
go run main.go --server_address :8082 &
go run main.go --server_address :8083