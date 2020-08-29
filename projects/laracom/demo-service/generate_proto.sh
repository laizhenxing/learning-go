#!/usr/bin/bash

cd proto/demo

protoc -I. --go_out=plugins=grpc:. demo.proto