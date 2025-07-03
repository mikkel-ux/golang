package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func UploadsWatchDog() {
	fmt.Println("Starting uploads watchdog...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v\n", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Printf("Event: %s\n", event)

			case err := <-watcher.Errors:
				fmt.Printf("Error: %v\n", err)
			}
		}
	}()

	if err := watcher.Add("./uploads"); err != nil {
		log.Fatalf("Failed to add directory to watcher: %v\n", err)
	}

	<-done
}
