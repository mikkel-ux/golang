package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var foo []string

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

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()
	uploadedTime := time.Now().UnixNano()
	uploadedTimeStr := fmt.Sprintf("%d", uploadedTime)
	name := strings.Split(fileHeader.Filename, ".")
	fileName := name[0] + filepath.Ext(fileHeader.Filename) + "___" + uploadedTimeStr
	dst, err := os.Create(fmt.Sprintf("./uploads/%s%s", fileName, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	foo = append(foo, dst.Name())

	/* fmt.Fprintf(w, "File uploaded successfully: %s", fileHeader.Filename) */
	response := ApiResponse{
		Message: fmt.Sprintf("File uploaded successfully: %s", fileHeader.Filename),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	/* for _, file := range foo {
		fmt.Fprintf(w, " File: %s, ", file)
	} */

}

func GetUploadsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	files, err := os.ReadDir("./uploads")
	if err != nil {
		http.Error(w, "Error reading uploads directory", http.StatusInternalServerError)
		return
	}

	var fileList []File
	for _, file := range files {
		fileList = append(fileList, SplitFileName(file))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fileList)

}

func FooTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	files, err := os.ReadDir("./uploads")
	if err != nil {
		http.Error(w, "Error reading uploads directory", http.StatusInternalServerError)
		return
	}

	/* foo = strings[0].Contains(files[0].Name(), "1753944929580453400") */
	for _, file := range files {
		if strings.Contains(file.Name(), "1753944929580453400") {
			log.Printf("Found file: %s\n", file.Name())
		}
	}

	/* log.Printf("Files in uploads directory: %v\n", files[0].Name()) */

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foo)

}

func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fileToDelete := ""
	found := false

	files, err := os.ReadDir("./uploads")
	if err != nil {
		http.Error(w, "Error reading uploads directory", http.StatusInternalServerError)
		return
	}

	fileID := r.URL.Query().Get("id")
	if fileID == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	for _, file := range files {
		if strings.Contains(file.Name(), fileID) {
			fileToDelete = strings.Join([]string{"./uploads", file.Name()}, "/")
			log.Printf("found file")
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	err = os.Remove(fileToDelete)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting file: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("File deleted successfully: %s\n", fileToDelete)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiResponse{Message: "File deleted successfully"})
}
