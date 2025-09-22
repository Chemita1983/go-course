package handlers

import (
	"fmt"
	"net/http"
)

// Handler -> Función que maneja una solicitud HTTP específica.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Página de inicio")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de")
}