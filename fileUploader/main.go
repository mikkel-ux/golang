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

func main() {
	content, err := fs.Sub(embeddedFiles, "fileUploaderFrontend/build")
	if err != nil {
		log.Fatalf("Failed to get subdirectory: %v\n", err)
	}
	http.HandleFunc("/api/upload", UploadHandler)

	fs := http.FileServer(http.FS(content))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path
		if filePath == "/" {
			filePath = "/index.html"
		}

		_, err := content.Open(path.Clean(filePath[1:]))
		if os.IsNotExist(err) {
			r.URL.Path = "/index.html"
		}

		fs.ServeHTTP(w, r)
	})

	if err := backend.CreateUploadsDir(); err != nil {
		panic(fmt.Sprintf("Failed to create uploads directory: %v\n", err))
	}

	go UploadsWatchDog()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
