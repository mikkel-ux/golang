package backend

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func UploadsWatchDog(broadcast chan File) {
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
					fileWithOutUploads := strings.Split(event.Name, "\\")
					fileNameSplited := strings.Split(fileWithOutUploads[1], "___")
					file := File{
						ID:   strings.Split(fileNameSplited[1], ".")[0],
						Name: fileNameSplited[0],
					}
					broadcast <- file
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
