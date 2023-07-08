init:
	go mod init skeleton

gen:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. proto/*/*.proto

.PHONY: init gen