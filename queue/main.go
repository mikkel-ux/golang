package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

var foo []string

const MAX_UPLOAD_SIZE = 10 * 1024 * 1024 // 10 MB

func uploadHandler(w http.ResponseWriter, r *http.Request) {
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

	checkDir, err := os.Stat("./uploads")
	if err != nil && os.IsNotExist(err) {
		checkDir = nil
	}

	if checkDir == nil {
		err = os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
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

	fmt.Fprintf(w, "File uploaded successfully: %s", fileHeader.Filename)

	for _, file := range foo {
		fmt.Fprintf(w, " File: %s, ", file)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	fs := http.FileServer(http.Dir("./fontend/static"))
	mux.Handle("/fontend/static/", http.StripPrefix("/fontend/static/", fs))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
