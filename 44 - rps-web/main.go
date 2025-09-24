package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

const PORT = ":8080"

func main() {

	// Router -> Componente que se encarga de dirigir las solicitudes a los controladores adecuados.
	router := http.NewServeMux()

	// Servir archivos estáticos (CSS, JS, imágenes) desde el directorio "static".
	fileStatic := http.FileServer(http.Dir("static/"))

	// Maneja las solicitudes a la ruta "/static/" y sirve los archivos estáticos.
	router.Handle("/static/", http.StripPrefix("/static/", fileStatic))

	// Configura las rutas y sus controladores.
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/newgame", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	log.Printf("Servidor escuchando en http://localhost%s\n", PORT)

	// Inicia el servidor web en el puerto especificado.
	log.Fatal(http.ListenAndServe(PORT, router))
}
