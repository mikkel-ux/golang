package main

import (
	"fileUploader/API"
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
	/* if err := backend.CreateUploadsDir(); err != nil {
		panic(fmt.Sprintf("Failed to create uploads directory: %v\n", err))
	}
	if err := backend.CreateUserDir(); err != nil {
		panic(fmt.Sprintf("Failed to create user directory: %v\n", err))
	} */
	if err := backend.CreateDir(); err != nil {
		panic(fmt.Sprintf("Failed to create directories: %v\n", err))
	}
	r := gin.Default()

	go backend.HandleMessages(broadcast)
	go backend.UploadsWatchDog(broadcast)

	api := r.Group("/api")
	{
		api.POST("/upload", API.UploadHandler)
		api.GET("/upload", API.GetUploadsHandler)
		api.GET("/upload/:id", API.DownloadFileHandler)
		api.DELETE("/upload/:id", API.DeleteFileHandler)
		api.GET("/video/:id", API.StreamVideoHandler)
		api.POST("/user", API.CreateUserHandler)
		api.POST("/login", API.LoginHandler)
		api.GET("/validate", API.ValidateTokenHandler)
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
