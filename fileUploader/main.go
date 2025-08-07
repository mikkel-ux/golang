package main

import (
	"fileUploader/backend"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

/* //go:embed all:fileUploaderFrontend/build */
/* var embeddedFiles embed.FS */

type ApiResponse struct {
	Message string `json:"message"`
}

var broadcast = make(chan backend.FileEvent)

func main() {
	if err := backend.CreateUploadsDir(); err != nil {
		panic(fmt.Sprintf("Failed to create uploads directory: %v\n", err))
	}
	r := gin.Default()

	go backend.HandleMessages(broadcast)
	go backend.UploadsWatchDog(broadcast)

	api := r.Group("/api")
	{
		api.POST("/upload", UploadHandler)
		api.GET("/upload", GetUploadsHandler)
		api.GET("/upload/:id", DownloadFileHandler)
		api.DELETE("/upload/:id", DeleteFileHandler)
	}

	r.GET("/ws", backend.SocketHandler)

	r.Static("/_app", "./fileUploaderFrontend/build/_app")

	r.NoRoute(func(c *gin.Context) {
		c.File("./fileUploaderFrontend/build/index.html")
	})

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
