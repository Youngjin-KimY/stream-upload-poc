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

	fmt.Println("Serve started at http://localhost:8888")

	err := http.ListenAndServe("0.0.0.0:8888", nil)
	// fmt.Println("https://localhost:8443")
	// err := http.ListenAndServeTLS(":8443", "localhost.crt", "localhost.key", nil)
	if err != nil {
		panic(err)
	}
}

func uploadStreamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======uploadStreamHandler======")
	// fmt.Println(r.Host)
	// fmt.Println(r.Method)
	// fmt.Println(r.Body)
	fmt.Printf("%s\n", r.URL.Query().Get("file_name"))
	filename := r.URL.Query().Get("file_name")

	if filename == "" {
		http.Error(w, "file_name is required", http.StatusBadRequest)
		return
	}

	filepath := buildFinalPath(filename)

	fmt.Printf("filepath: %s\n", filepath)

	dst, err := os.Create(filepath)

	if err != nil {
		msg := fmt.Sprintf("error: %s\n", err.Error())
		fmt.Println(msg)
		http.Error(w, "Fail to upload1", http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, r.Body)

	if err != nil {
		msg := fmt.Sprintf("error: %s\n", err.Error())
		fmt.Println(msg)
		http.Error(w, "Fail to upload2", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("upload completed from server"))
}

func buildFinalPath(fileName string) string {
	today := time.Now().Format("200601")
	pwd, _ := os.Getwd()
	// fmt.Printf("chdir: %s\n", pwd)
	dateFolder := filepath.Join(pwd, storageRoot, today)
	_, err := os.Stat(dateFolder)
	if os.IsNotExist(err) {
		os.Mkdir(dateFolder, 0755)
	}

	return filepath.Join(pwd, storageRoot, today, fileName)
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
