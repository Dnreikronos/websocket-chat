build:
	@go build -o bin/websocket-chat cmd/main.go

run: build
	@.bin/websocket-chat cmd/main.go

test:
	@go test -v ./..
