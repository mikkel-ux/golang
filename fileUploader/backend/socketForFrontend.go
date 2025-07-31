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
	log.Printf("client key %v\n", conn)

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			log.Printf("Failed to read message from client: %v\n", err)
			clients.Lock()
			delete(clients.m, conn)
			clients.Unlock()
			break
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

}
