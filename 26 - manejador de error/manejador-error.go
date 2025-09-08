package main

import (
	"fmt"
	"log"

	"github.com/Chemita1983/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(3) // Establece la bandera de formato en 0 (con 0 no se muestran), las banderas de formato determinan información adicional, como la fecha y la hora

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
