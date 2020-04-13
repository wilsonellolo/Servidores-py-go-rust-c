package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var templatesDir = os.Getenv("static")

func pagina(w http.ResponseWriter, r *http.Request) {

	tmplPath := filepath.Join("static", "hola.html")

	tmpl := template.Must(template.ParseFiles(tmplPath))

	data := "La Chartreuse"

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", pagina)
	http.ListenAndServe(":9000", nil)
}
