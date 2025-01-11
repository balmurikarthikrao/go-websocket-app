# Go WebSocket Notification Application

This project is a WebSocket application built in Go that allows sending notifications to users based on their unique IDs. 

## Project Structure

```
go-websocket-app
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── websocket.go # Handles WebSocket connections
│   ├── services
│   │   └── notification.go # Sends notifications to users
│   └── models
│       └── user.go      # Defines the User model
├── go.mod               # Module definition and dependencies
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-websocket-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## Usage

- Connect to the WebSocket server using a WebSocket client.
- Send a message containing the user ID to establish a connection.
- Use the `SendNotification` function to send notifications to the connected user.

## Example

To send a notification to a user with ID `123`, you would call:
```go
SendNotification("123", "Your notification message")
```

## License

This project is licensed under the MIT License.