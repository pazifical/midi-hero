package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/pazifical/midi-hero/internal/clonehero"
	"github.com/pazifical/midi-hero/internal/midi"
)

//go:embed static templates
var embedFS embed.FS

var port uint = 12468

var chartDirectory = "charts"

var templates *template.Template

func init() {
	t, err := template.ParseFS(embedFS, "templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	templates = t

	os.Mkdir(chartDirectory, 0750)
}

func main() {
	http.Handle("GET /static/", http.FileServerFS(embedFS))

	http.HandleFunc("GET /", serveIndex)
	http.HandleFunc("POST /api/process", processMidi)

	err := openDefaultWebBrowser(fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		log.Println(err)
	}

	err = http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
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

	fhs := r.MultipartForm.File["files"]
	for _, fileHeader := range fhs {
		file, err := fileHeader.Open()
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

	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func openDefaultWebBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
