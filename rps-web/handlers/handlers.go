package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	TEMPLATE_DIR  = "templates/"
	TEMPLATE_BASE = TEMPLATE_DIR + "base.html"
)

// Handler -> Función que maneja una solicitud HTTP específica.
func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "newgame.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "game.html", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	// Parsear las plantillas HTML.
	tmpl := template.Must(template.ParseFiles(TEMPLATE_BASE, TEMPLATE_DIR+page))

	// Renderizar la plantilla con datos (aquí no se pasan datos, por simplicidad).
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al ejecutar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
