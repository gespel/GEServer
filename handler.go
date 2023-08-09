package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/download/")
	filePath := path.Join("tracks", filename) // Passe den Pfad zum Ordner an

	http.ServeFile(w, r, filePath)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := getTracks()

	type FileWithLink struct {
		Name string
		Link string
	}
	var filesWithLinks []FileWithLink
	for _, filename := range files {
		link := "/download/" + filename
		filesWithLinks = append(filesWithLinks, FileWithLink{Name: filename, Link: link})
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, filesWithLinks)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("upload.html"))
	tmpl.Execute(w, nil)
}

func uploaderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Die hochgeladene Datei wird im Form-Attribut "file" gesendet
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, "Error reading files", http.StatusInternalServerError)
		return
	}

	// Schleife über die hochgeladenen Dateien
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Error processing files", http.StatusInternalServerError)
			return
		}

		// Datei auf dem Server speichern
		file, err := os.Create("tracks/" + part.FileName())
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		io.Copy(file, part)
	}

	fmt.Fprint(w, "Dateien wurden erfolgreich hochgeladen!")
}
