#!/usr/bin/bash

cd model/protos/prods

protoc --micro_out=. --go_out=. prods.proto
protoc-go-inject-tag -input=./prods.pb.go
cd ../../../