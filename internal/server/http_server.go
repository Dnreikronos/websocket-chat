package server

import (
	"log"
	"net/http"

	"github.com/Dnreikronos/websocket-chat/internal/chat"
)

func StartHTTPServer(addr string) {
	manager := chat.NewManager()
	go manager.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.HandleWebSocketConnection(manager, w, r)
	})

	log.Printf("HTTP server started on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
