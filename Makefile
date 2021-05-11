VERSION ?= test-$(shell git log --pretty=format:'%h' -n 1)

all: build

build:
	@go build -ldflags "-s -w" -o msatodo ./cmd/msatodo/*.go

docker:
	@docker build -t msatodo:$(VERSION) .

protoc:
	@protoc -I=. \
		--go_out . \
		--go_opt paths=source_relative \
		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		`find ./pkg -iname "*.proto"`
