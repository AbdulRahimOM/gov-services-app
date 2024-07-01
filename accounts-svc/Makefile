.PHONY: 
	server \
	running \
	running-local \
	running-in-background

server:
	go run cmd/main.go

running-local:
	CompileDaemon  -build="go build -o cmd/main cmd/main.go"  -command="./cmd/main"

running:
	CompileDaemon -directory=. -directory=../internal -directory=./../shared  -build="go build -o cmd/main cmd/main.go"  -command="./cmd/main"

running-in-background:
	CompileDaemon -directory=. -directory=../internal -directory=./../shared  -build="go build -o cmd/main cmd/main.go"  -command="./cmd/main"&