package main

import (
	"fmt"
	"io"
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

	// fmt.Println("Serve started at http://localhost:8080")

	// err := http.ListenAndServe(":8080", nil)
	fmt.Println("https://localhost:8443")
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	if err != nil {
		panic(err)
	}
}

func uploadStreamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======uploadStreamHandler======")
	// fmt.Println(r.Host)
	// fmt.Println(r.Method)
	// fmt.Println(r.Body)

	filename := r.URL.Query().Get("file_name")
	if filename == "" {
		http.Error(w, "file_name is required", http.StatusBadRequest)
		return
	}

	filepath := buildFinalPath(filename)
	// fmt.Println(filepath)
	dst, err := os.Create(filepath)

	if err != nil {
		http.Error(w, "Fail to upload1", http.StatusInternalServerError)
	}

	defer dst.Close()

	_, err = io.Copy(dst, r.Body)
	if err != nil {
		http.Error(w, "Fail to upload2", http.StatusInternalServerError)
	}

	w.Write([]byte("upload completed from server"))
}

func buildFinalPath(fileName string) string {
	today := time.Now().Format("200601")
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
