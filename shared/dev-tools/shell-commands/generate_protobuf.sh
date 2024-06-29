#!/bin/zsh

# cd shared
# echo "Current directory: $(pwd)"

protoc --go_out=. --go-grpc_out=. ./**/*.proto

echo "Protobuf files compilation attempt complete."

cd -