package main

import (
	"log"
	"os"
)

func main() {

	/* Podemos enviar el registro a un archivo, para posterior procesamiento, para ocultar información
	al usuario final
	*/

	log.SetPrefix("main(): ") // Agrega un prefijo a los mensajes generados

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // Creamos un archivo que ira añadiendo mensajes y de solo escritura,  además con permisos de escritura y lectura para el usuario (0644)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer file.Close()

	log.SetOutput(file) // Enviamos el archivo donde se van a registrar los errores
	log.Print("Soy un mensaje de log!!")

}

func showLogs() {
	log.Println("Es es un mensaje de log")
	log.Println("Este es otro mensaje de log")
	log.Panic("Mensaje de log panic") // Detiene la ejecución del programa y muestra pila de errores
	log.Fatal("Mensaje de log fatal") // Detiene la ejecución del programa

}
