/*
Como el paquete aun no esta instalado ni desplegado hay que reemplazar el valor por el que tenemos en local

➜  hello git:(main) ✗ go mod edit -replace github.com/Chemita1983/greetings=../../24\ -\ crear-modulo/greetings
➜  hello git:(main) ✗ go mod tidy
go: found github.com/Chemita1983/greetings in github.com/Chemita1983/greetings v0.0.0-00010101000000-000000000000
*/

package main

import (
	"fmt"

	"github.com/Chemita1983/greetings"
)

func main() {
	message, _ := greetings.Hello("Chema")
	fmt.Println(message)
}
