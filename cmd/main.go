package main

import (
	"log"

	"github.com/Dnreikronos/websocket-chat/internal/server"
)



func main() {
	log.Println("Starting chat server")
	server.StartHTTPServer(":8080")
}
