package backend

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type File struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

type FileEvent struct {
	File           File   `json:"file"`
	FileWasRemoved string `json:"fileWasRemoved"`
	ClientsCount   *int   `json:"clientsCount"`
}

func UploadsWatchDog(broadcast chan FileEvent) {
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
					fmt.Println("File watcher closed")
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					fileWithOutUploads := strings.Split(event.Name, "\\")
					fmt.Println("File created1:", fileWithOutUploads[1])
					fileNameSplited := strings.Split(fileWithOutUploads[1], "___")
					file := File{
						ID:        strings.Split(fileNameSplited[2], ".")[0],
						Name:      fileNameSplited[0],
						Extension: strings.Split(fileNameSplited[2], ".")[1],
					}
					fileEvent := FileEvent{
						File:           file,
						FileWasRemoved: "",
						ClientsCount:   nil,
					}
					broadcast <- fileEvent
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fileWithOutUploads := strings.Split(event.Name, "\\")
					fileNameSplited := strings.Split(fileWithOutUploads[1], "___")
					FileEvent := FileEvent{
						File:           File{},
						FileWasRemoved: strings.Split(fileNameSplited[2], ".")[0],
						ClientsCount:   nil,
					}
					broadcast <- FileEvent
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
