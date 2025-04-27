package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const storageRoot = "./hdd"
const port = 8080

func main() {
	if !directoryExists(storageRoot) {
		err := os.Mkdir(storageRoot, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Failed to create storage root: %v", err))
		}
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upload_stream", uploadStreamHandler)

	fmt.Println("Serve started at http://localhost:80")
	http.ListenAndServe(":80", nil)
}

func uploadStreamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploadStreamHandler")
	fmt.Sprintln(r.Host)
	fmt.Sprintln(r.Method)
	fmt.Sprintln(r.Body)

	filename := r.URL.Query().Get("file_name")
	if filename == "" {
		http.Error(w, "file_name is required", http.StatusBadRequest)
		return
	}

	filepath := buildFinalPath(filename)
	fmt.Sprintf("%s", filepath)
}

func buildFinalPath(fileName string) string {
	today := time.Now().Format("20061111")
	return filepath.Join(storageRoot, today, fileName)
}

func directoryExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
