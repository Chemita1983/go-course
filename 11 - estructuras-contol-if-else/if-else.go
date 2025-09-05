// Verificamos la hora y dependiendo de ella, decimos si es mañana, tarde o noche.

package main

import (
	"fmt"
	"time"
)

func main() {

	if tiempo := time.Now(); tiempo.Hour() < 15 { // Podemos declarar una variable dentro del if, solo se usa en ese if y sus anexionados
		fmt.Println("Es por la mañana")
	} else if tiempo.Hour() < 21 {
		fmt.Println("Es por la tarde")
	} else {
		fmt.Println("Es por la noche")
	}

}
