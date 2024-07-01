.PHONY: 
	proto \
	watching-proto

proto:	
	protoc --go_out=. --go-grpc_out=. ./**/*.proto

watching-proto:
	CompileDaemon  \
	-directory=. \
	-build="true" \
	-command="./dev-tools/shell-commands/generate_protobuf.sh"  \
	-pattern=".*\.proto$\" 

watching-proto-in-background:
	CompileDaemon  \
	-directory=. \
	-build="true" \
	-command="./dev-tools/shell-commands/generate_protobuf.sh"  \
	-pattern=".*\.proto$\" &