package main

import (
	"go-websocket-app/internal/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ws", handlers.HandleWebSocket)

	serverAddr := "localhost:8080"
	log.Printf("Starting WebSocket server at %s\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
