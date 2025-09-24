package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
	"text/template"
)

const (
	TEMPLATE_DIR  = "templates/"
	TEMPLATE_BASE = TEMPLATE_DIR + "base.html"
)

type Player struct {
	Name string
}

var player Player

// Handler -> Función que maneja una solicitud HTTP específica.
func Index(w http.ResponseWriter, r *http.Request) {
	resetValues()
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	resetValues()
	renderTemplate(w, "newgame.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}

	// Si el nombre del jugador está vacío, redirigir a /newgame.
	if player.Name == "" {
		http.Redirect(w, r, "/newgame", http.StatusFound)
		return
	}

	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	resetValues()
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

// Reiniciar el valor del jugador
func resetValues() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
