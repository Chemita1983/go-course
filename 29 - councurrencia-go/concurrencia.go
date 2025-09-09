/*
  La concurrencia en go es la ejecucion de multiples tareas o Goruntines.
  Goruntimes son similares a hilos.

  Los canales son una estructura que permite mandar y recibir canales entre goruntimes, actuando como un conducto a
  través del cual fluye la información

  	/* Canales
	channel := make(chan int)
	channel <- 15 // enviar datos al canal
	value := <- channel // recibir datos del canal
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	ch := make(chan string)
	start := time.Now()

	apis := []string{
		"https://management.azure.com/",
		"https://dev.azure.com/",
		"https://api.github.com/",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkAPI(api, ch) // Aplica la concurrencia en go.
	}

	// Para mostrar los resultados del canal
	for i := 0; i < len(apis); i++ {
		fmt.Printf(<-ch) // Mostramos datos del canal
	}

	elapsed := time.Since(start) // Calcula el tiempo que tarda en ejecutarse nuestra aplicacion
	fmt.Printf("Finish: %v seconds\n", elapsed.Seconds())

}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: %s is Down\n", api) // Mandamos el resultado al canal
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is OK\n", api)

}
