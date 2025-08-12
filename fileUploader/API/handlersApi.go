package API

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
)

const MAX_UPLOAD_SIZE = 10 * 1024 * 1024

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func SplitFileName(file os.DirEntry) File {
	fileNameSplited := strings.Split(file.Name(), "___")
	return File{
		ID:   strings.Split(fileNameSplited[1], ".")[0],
		Name: fileNameSplited[0],
	}
}

func findFileByID(fileID string) (string, bool, error) {
	files, err := os.ReadDir("./uploads")
	if err != nil {
		return "", false, fmt.Errorf("error reading uploads directory: %v", err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), fileID) {
			return strings.Join([]string{"./uploads", file.Name()}, "/"), true, nil
		}
	}
	return "", false, fmt.Errorf("file with ID %s not found", fileID)
}

func UploadHandler(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MAX_UPLOAD_SIZE)
	if err := c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The uploaded file is too big"})
		return
	}

	file, _ := c.FormFile("file")
	uploadedTime := time.Now().UnixNano()
	uploadedTimeStr := fmt.Sprintf("%d", uploadedTime)
	name := strings.Split(file.Filename, ".")
	newFileName := name[0] + filepath.Ext(file.Filename) + "___" + uploadedTimeStr + filepath.Ext(file.Filename)
	c.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s", newFileName))

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File uploaded successfully: %s", newFileName)})
}

func GetUploadsHandler(c *gin.Context) {
	files, err := os.ReadDir("./uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading uploads directory"})
		return
	}

	var fileList []File
	for _, file := range files {
		fileList = append(fileList, SplitFileName(file))
	}
	c.JSON(http.StatusOK, fileList)
}

func DeleteFileHandler(c *gin.Context) {
	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	fileToDelete, found, err := findFileByID(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding file: %v", err)})
		return
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	if os.Remove(fileToDelete) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting file: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

func DownloadFileHandler(c *gin.Context) {
	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	filePath, found, err := findFileByID(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding file: %v", err)})
		return
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting file info: %v", err)})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	c.FileAttachment(filePath, fileInfo.Name())
}

func StreamVideoHandler(c *gin.Context) {
	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	filePath, found, err := findFileByID(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding file: %v", err)})
		return
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error opening file: %v", err)})
		return
	}
	defer file.Close()

	c.Header("Content-Type", "video/mp4")
	c.Header("Accept-Ranges", "bytes")
	c.Header("Cache-Control", "no-cache")
	c.File(filePath)
}

func TestHandler(c *gin.Context) {
	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	filePath, found, err := findFileByID(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding file: %v", err)})
		return
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	head := make([]byte, 512)
	/* file.Read(head) */
	n, err := file.Read(head)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	head = head[:n]

	if filetype.IsVideo(head) {
		fmt.Println("File is an video")
	} else if filetype.IsImage(head) {
		fmt.Println("file is an image")
	} else if filetype.IsAudio(head) {
		fmt.Println("file is an audio")
	} else if filetype.IsDocument(head) {
		fmt.Println("file is a document")
	} else if filetype.IsArchive(head) {
		fmt.Println("file is an archive")
	} else if IsTextFile(file.Name(), head) {
		fmt.Println("file is a text file")
	} else {
		fmt.Println("file is an unknown type")
	}
}

func IsTextFile(filename string, content []byte) bool {
	// Check extension first (fast path)
	if filepath.Ext(filename) == ".txt" {
		return true
	}

	// Verify content when extension missing/wrong
	return false
}
