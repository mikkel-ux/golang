package backend

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var clients = struct {
	sync.RWMutex
	m map[*websocket.Conn]bool
}{m: make(map[*websocket.Conn]bool)}

/* var broadcast = make(chan string) */

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

func UploadsWatchDogTest() {
	fmt.Println("Starting uploads watchdog...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v\n", err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					file := strings.Split(event.Name, "\\")
					log.Printf("File uploaded: %s\n", file[1])
					/* broadcast <- event.Name */
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("Error: %v\n", err)
			}
		}
	}()

	if err := watcher.Add("./uploads"); err != nil {
		log.Fatalf("Failed to add directory to watcher: %v\n", err)
	}

	<-make(chan bool)
}
