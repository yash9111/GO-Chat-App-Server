package routes

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Conn   *websocket.Conn
	UserID string
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	rooms = make(map[string][]*Connection) // Map room ID to list of connections
	mu    sync.Mutex                       // Mutex to manage concurrent access to rooms map
)

// ServeWs handles WebSocket requests and assigns them to a specific room.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room") // Use query param "room" to identify the chat room
	userID := r.URL.Query().Get("user") // Identify user (optional, for tracking)

	if roomID == "" || userID == "" {
		http.Error(w, "room and user are required query parameters", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	client := &Connection{Conn: conn, UserID: userID}
	addClientToRoom(roomID, client)
	defer removeClientFromRoom(roomID, client)

	fmt.Printf("User %s joined room %s\n", userID, roomID)

	for {
		// Read message from the WebSocket
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Message from %s in room %s: %s\n", userID, roomID, msg)

		// Broadcast the message to other users in the same room
		broadcastToRoom(roomID, msg, userID)
	}
}

// Add client to room
func addClientToRoom(roomID string, client *Connection) {
	mu.Lock()
	defer mu.Unlock()
	rooms[roomID] = append(rooms[roomID], client)
}

// Remove client from room
func removeClientFromRoom(roomID string, client *Connection) {
	mu.Lock()
	defer mu.Unlock()
	conns := rooms[roomID]
	for i, c := range conns {
		if c == client {
			rooms[roomID] = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	if len(rooms[roomID]) == 0 {
		delete(rooms, roomID) // Remove empty room
	}
}

// Broadcast message to all clients in the room except sender
func broadcastToRoom(roomID string, message []byte, senderID string) {
	mu.Lock()
	defer mu.Unlock()

	for _, client := range rooms[roomID] {
		if client.UserID != senderID { // Skip sender
			err := client.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println("Broadcast error:", err)
			}
		}
	}
}
