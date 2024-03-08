#!/usr/bin/env bash

find . -type d  -exec rm -rf {} \;

protoc --go_out . --go-grpc_out . --go_opt module=iceline-hosting.com/backend/proto --go-grpc_opt module=iceline-hosting.com/backend/proto ./*.proto