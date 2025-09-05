package API

import (
	"fileUploader/models"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const MAX_UPLOAD_SIZE = 10 * 1024 * 1024

/* type File struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
	FileType  string `json:"fileType"`
} */

func SplitFileName(file os.DirEntry) models.File {
	fileNameSplited := strings.Split(file.Name(), "___")
	return models.File{
		ID:        strings.Split(fileNameSplited[2], ".")[0],
		Name:      fileNameSplited[0],
		Extension: strings.Split(fileNameSplited[2], ".")[1],
		FileType:  fileNameSplited[1],
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

/*
c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MAX_UPLOAD_SIZE)

	if err := c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The uploaded file is too big"})
		return
	}
*/
//TODO add file type to file name
func UploadHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if err := ValidateTokenFunc(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	file, _ := c.FormFile("file")
	uploadedTime := time.Now().UnixNano()
	uploadedTimeStr := fmt.Sprintf("%d", uploadedTime)
	fileName := strings.ReplaceAll(file.Filename, " ", "_")
	contentType := strings.ReplaceAll(file.Header.Get("Content-Type"), "/", "__")
	fmt.Println(contentType)
	newFileName := fileName + "___" + contentType + "___" + uploadedTimeStr + filepath.Ext(file.Filename)

	fmt.Println("file.Filename:", file.Header.Get("Content-Type"))

	isVideo := strings.HasPrefix(file.Header.Get("Content-Type"), "video/")
	if isVideo {
		if err := MoveMetaDataAndFileToUploads(file, newFileName, c); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s", newFileName))
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File uploaded successfully: %s", newFileName)})
}

func MoveMetaDataAndFileToUploads(file *multipart.FileHeader, newFileName string, c *gin.Context) error {
	tempPath := fmt.Sprintf("./temp/%s", newFileName)
	uploadPath := fmt.Sprintf("./uploads/%s", newFileName)
	c.SaveUploadedFile(file, tempPath)

	cmd := exec.Command("ffmpeg", "-i", tempPath,
		"-c:v", "copy", "-c:a", "copy", "-movflags", "+faststart",
		uploadPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error moving file to uploads: %v", err)
	}
	os.Remove(tempPath)

	return nil
}

func GetUploadsHandler(c *gin.Context) {
	files, err := os.ReadDir("./uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading uploads directory"})
		return
	}

	var fileList []models.File
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
	token := c.GetHeader("Authorization")
	if err := ValidateTokenFunc(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
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

	/* c.Header("Content-Type", "video/mp4")
	c.Header("Accept-Ranges", "bytes")
	c.Header("Cache-Control", "no-cache")
	c.File(filePath) */

	stat, err := file.Stat()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Could not obtain file size"})
		return
	}
	http.ServeContent(c.Writer, c.Request, file.Name(), stat.ModTime(), file)
}
