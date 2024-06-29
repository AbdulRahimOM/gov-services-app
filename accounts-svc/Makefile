.PHONY: 
	proto \
	server \
	running \
	running-local \
	running-in-background
	
proto:
	protoc --go_out=. --go-grpc_out=. ./internal/infrastructure/pb/protobuf/*.proto

server:
	go run cmd/main.go

running-local:
	CompileDaemon  -build="go build -o cmd/main cmd/main.go"  -command="./cmd/main"

running:
	CompileDaemon -directory=. -directory=../internal -directory=./../shared  -build="go build -o cmd/main cmd/main.go"  -command="./cmd/main"

running-in-background:
	CompileDaemon -directory=. -directory=../internal -directory=./../shared  -build="go build -o cmd/main cmd/main.go"  -command="./cmd/main"&