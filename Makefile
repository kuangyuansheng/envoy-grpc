all:
	go build -o rpcSrv ./cmd/rpc/*.go
	go build -o rpcCli ./cmd/client/*.go
