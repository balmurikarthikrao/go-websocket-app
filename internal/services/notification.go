package services

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	Conn *websocket.Conn
}

type User struct {
	ID         string
	Connection *WebSocketConnection // Assume WebSocketConnection is defined elsewhere
}

var (
	users      = make(map[string]*User)
	usersMutex sync.Mutex
)

func (ws *WebSocketConnection) Send(message string) error {
	return ws.Conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func SendNotification(userID string, message string) error {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, exists := users[userID]
	if !exists {
		return fmt.Errorf("user with ID %s not found", userID)
	}

	err := user.Connection.Send(message) // Assume Send method is defined on WebSocketConnection
	if err != nil {
		return fmt.Errorf("failed to send notification to user %s: %v", userID, err)
	}

	return nil
}

func RegisterUser(userID string, connection *WebSocketConnection) {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	users[userID] = &User{
		ID:         userID,
		Connection: connection,
	}
}

func UnregisterUser(userID string) {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	delete(users, userID)
}
