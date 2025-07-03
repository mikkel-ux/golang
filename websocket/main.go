package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string `json:"id"`
	Conn *websocket.Conn
}

type Message struct {
	TargetID string `json:"targetId"`
	Content  string `json:"content"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		/* origin := r.Header.Get("Origin")
		return origin == "http://localhost:8080" */
		return true // Allow all origins for simplicity; adjust as needed
	},
}

var clients = make(map[string]*websocket.Conn)

/* var broadcast = make(chan []byte) */
var mutex = &sync.Mutex{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("id")
	if clientID == "" {
		http.Error(w, "Client ID is required", http.StatusBadRequest)
		fmt.Println("Client ID is required")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	if _, ok := clients[clientID]; ok {
		mutex.Unlock()
		fmt.Printf("Client with ID %s already connected. Connection rejected.\n", clientID)
		conn.WriteMessage(websocket.TextMessage, []byte("Connection rejected: Client ID already in use."))
		return
	}
	clients[clientID] = conn
	mutex.Unlock()

	fmt.Printf("Client with ID %s connected.\n", clientID)

	defer func() {
		mutex.Lock()
		delete(clients, clientID)
		mutex.Unlock()
		fmt.Printf("Client with ID %s disconnected.\n", clientID)
	}()

	handleMessages(clientID, conn)
}

func handleMessages(clientID string, conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var msg Message
		if err := json.Unmarshal(p, &msg); err != nil {
			fmt.Printf("Error unmarshalling message from %s: %v\n", clientID, err)
			continue
		}

		mutex.Lock()
		targetConn, found := clients[msg.TargetID]
		mutex.Unlock()

		if found {
			fmt.Printf("Relaying message from %s to %s\n", clientID, msg.TargetID)
			err := targetConn.WriteMessage(websocket.TextMessage, p)
			if err != nil {
				fmt.Printf("Error sending message to %s: %v\n", msg.TargetID, err)
			}
		} else {
			fmt.Printf("Client %s tried to send message to non-existent client %s\n", clientID, msg.TargetID)
			conn.WriteMessage(websocket.TextMessage, []byte("Error: Target user not found."))
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket server started on :8080")
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		fmt.Println("Error starting server:", er)
	}
}
