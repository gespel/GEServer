package main

import (
	"fmt"
	"net/http"
	"os"
)

func initFilesystem() {
	if _, err := os.Stat("tracks"); os.IsNotExist(err) {
		err := os.Mkdir("tracks", os.ModePerm)
		if err != nil {
			return
		}
	}
}

func main() {
	initFilesystem()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/uploader", uploaderHandler)
	http.HandleFunc("/download/", downloadHandler)

	port := ":8080"
	fmt.Println("Server läuft auf http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return
	}
}
