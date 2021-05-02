
all: build

build:
	@go build -a -ldflags '-extldflags "-static"' -o prototodo ./cmd/prototodo/*.go

protoc:
	@protoc -I=. \
		--go_out . \
		--go_opt paths=source_relative \
		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		protos/v1/**/**.proto