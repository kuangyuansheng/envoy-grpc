# alpine: CGO_ENABLED=0 go build
all:
	go build -o rpc_srv ./cmd/rpc/*.go
	go build -o rpc_check ./cmd/client/*.go
