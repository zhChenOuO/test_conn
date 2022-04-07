p=$(shell pwd)

gen.proto:
	protoc -I$(CURDIR) --gofast_out=plugins=grpc:. $(CURDIR)/proto/*.proto

server:
	PROJ_DIR=$p CONFIG_PATH=deployment/config go run main.go server

client:
	PROJ_DIR=$p CONFIG_PATH=deployment/config go run main.go client
