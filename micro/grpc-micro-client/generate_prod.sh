#!/usr/bin/bash

echo "generate protobuf go and micro files"

cd proto/prod
# prodModel
protoc --micro_out=. --go_out=. models.proto
## prodService
protoc --micro_out=. --go_out=. prodService.proto
# inject-tag
protoc-go-inject-tag -input=./models.pb.go

cd ../../../

echo "generate finish!"