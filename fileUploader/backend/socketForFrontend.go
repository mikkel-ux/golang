package backend

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = struct {
	sync.RWMutex
	m map[*websocket.Conn]bool
}{m: make(map[*websocket.Conn]bool)}

var firstFrontendKey *websocket.Conn

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v\n", err)
		return
	}

	defer conn.Close()

	clients.Lock()
	clients.m[conn] = true
	clients.Unlock()

	log.Println("Client connected")
	HandleMessageToSpecificClient(len(clients.m))

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message from client: %v\n", err)
			conn.Close()
			clients.Lock()
			delete(clients.m, conn)
			clients.Unlock()
			if firstFrontendKey == conn {
				firstFrontendKey = nil
			} else {
				HandleMessageToSpecificClient(len(clients.m))
			}
			break
		}
		log.Printf("Received message: %s\n", message)
		if string(message) == "serverFrontend" {
			if firstFrontendKey == nil {
				firstFrontendKey = conn
				HandleMessageToSpecificClient(len(clients.m))
			}
		}
	}
}

func HandleMessages(broadcast chan File) {
	for {
		msg := <-broadcast
		data, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Failed to marshal message: %v\n", err)
			continue
		}
		log.Printf("clients %d\n", len(clients.m))
		for client := range clients.m {
			err := client.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Printf("Failed to write message to client: %v\n", err)
				client.Close()
			}
		}
		log.Printf("Broadcasted message: %s\n", msg)
		clients.RLock()
	}
}

func HandleMessageToSpecificClient(clientCount int) {
	data, err := json.Marshal(clientCount)
	if err != nil {
		log.Printf("Failed to marshal client count: %v\n", err)
		return
	}
	if firstFrontendKey != nil {
		err := firstFrontendKey.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Printf("Failed to write message to first client: %v\n", err)
			firstFrontendKey.Close()
		}
	}
}
