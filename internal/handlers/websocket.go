package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type User struct {
	ID         string
	Connection *websocket.Conn
}

var (
	users      = make(map[string]*User)
	usersMutex sync.Mutex
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	var userID = r.PostFormValue("userID")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	usersMutex.Lock()
	users[userID] = &User{ID: userID, Connection: conn}
	usersMutex.Unlock()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Received message from user %s: %s\n", userID, message)
	}

	usersMutex.Lock()
	delete(users, userID)
	usersMutex.Unlock()
}
