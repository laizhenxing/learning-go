#!/usr/bin/bash

cd proto/hellomicro

protoc --micro_out=. --go_out=. hellomicro.proto
protoc-go-inject-tag -input=./hellomicro.pb.go
echo "generate hellomicro~"

cd ../user/

protoc --go_out=. user.proto
protoc --micro_out=. --go_out=. userService.proto
protoc-go-inject-tag -input=./user.pb.go
protoc-go-inject-tag -input=./userService.pb.go
echo "generate user~"

cd ../../


