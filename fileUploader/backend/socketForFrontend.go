package backend

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var clients = struct {
	sync.RWMutex
	m map[*websocket.Conn]bool
}{m: make(map[*websocket.Conn]bool)}

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

func HandleMessages(broadcast chan string) {
	for {
		msg := <-broadcast
		for client := range clients.m {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("Failed to write message to client: %v\n", err)
				client.Close()
			}
		}
		log.Printf("Broadcasted message: %s\n", msg)
		clients.RLock()
	}
}
