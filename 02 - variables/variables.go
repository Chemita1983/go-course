/* Declaracion de variables

Se pueden declarar variables del mismo tipo en una linea

Si no se usa una variable, el compilador dará error.

Las variables se pueden declarar de varias maneras:

		var firstName, lastName string
		var age int

		var (
			firstName, lastName string
			age int
		)

		var (
			firstName = "Chema"
			lastName  = "J.V"
			age       = 42

	Cuando inicializamos los datos de una variable no es necesario
	indicarle el tipo de dato.

	También se puede declarar todo en una sola linea.

		var firstName, lastName, age = "Chema", "J.V", 42

	Estas variables podriamos declararlas fuera de la funcion main

	Podemos usar := para declarar variables (solo declarar, no modificar)
	sin necesidad de escribir var, pero solo dentro de las funciones.
	(Esta es la forma que mas se va a usar)

		firstName, lastName, age := "Chema", "J.V", 42
*/

package main

import (
	"fmt"
)

func main() {

	firstName, lastName, age := "Chema", "J.V", 42

	fmt.Println(firstName, lastName, age)
}
