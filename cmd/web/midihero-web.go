package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pazifical/midi-hero/internal/clonehero"
	"github.com/pazifical/midi-hero/internal/midi"
)

//go:embed static templates
var embedFS embed.FS

var port uint = 8080

var chartDirectory = "charts"

var templates *template.Template

func init() {
	t, err := template.ParseFS(embedFS, "templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	templates = t

	err = os.Mkdir(chartDirectory, 0750)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.Handle("GET /static/", http.FileServerFS(embedFS))

	http.HandleFunc("GET /", serveIndex)
	http.HandleFunc("POST /api/process", processMidi)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func processMidi(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	chart, err := midi.ImportFromReader(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("%s.chart", fileHeader.Filename)
	filePath := filepath.Join(chartDirectory, fileName)
	err = clonehero.WriteToFile(chart, fmt.Sprintf("%s.chart", filePath))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
