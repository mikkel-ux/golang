package main

import (
	"embed"
	"fileUploader/backend"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
)

//go:embed all:fileUploaderFrontend/build
var embeddedFiles embed.FS

type ApiResponse struct {
	Message string `json:"message"`
}

var broadcast = make(chan backend.File)

func main() {
	if err := backend.CreateUploadsDir(); err != nil {
		panic(fmt.Sprintf("Failed to create uploads directory: %v\n", err))
	}

	go backend.HandleMessages(broadcast)
	go backend.UploadsWatchDog(broadcast)

	http.HandleFunc("/api/upload", UploadHandler)
	http.HandleFunc("/api/files", GetUploadsHandler)
	http.HandleFunc("/api/delete/{id}", DeleteFileHandler)
	http.HandleFunc("/api/test", FooTest)
	http.HandleFunc("/ws", backend.SocketHandler)

	svelteFS, err := fs.Sub(embeddedFiles, "fileUploaderFrontend/build")
	if err != nil {
		log.Fatalf("Failed to get subdirectory: %v\n", err)
	}
	fileServer := http.FileServer(http.FS(svelteFS))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path

		cleandPath := path.Clean(filePath)

		if filePath == "/" {
			http.ServeFileFS(w, r, svelteFS, "index.html")
			return
		}

		_, err := svelteFS.Open(cleandPath[1:])
		if os.IsNotExist(err) {
			http.ServeFileFS(w, r, svelteFS, "index.html")
			return
		} else if err != nil {
			log.Printf("Error opening file %s: %v", filePath, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fileServer.ServeHTTP(w, r)

	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
