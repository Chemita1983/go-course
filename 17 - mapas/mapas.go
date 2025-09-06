// Estructura de datos que guarda datos con clave - valor

package main

import "fmt"

func main() {

	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors)         // Imprime todos los elementos
	fmt.Println(colors["rojo"]) // Imprime el elemento que le indiquemos en la key

	// Agregamos nuevo elemento
	colors["negro"] = "#000000"
	fmt.Println(colors) // No respeta el orden, lo guarda por orden alfabetico.

	valor, exist := colors["rojo"] // Podemos devolver una verificación que nos indica si existe o no el elemento
	fmt.Println(valor, exist)

	valorNoExiste, exist := colors["blanco"] // Podemos devolver una verificación que nos indica si existe o no el elemento
	fmt.Println(valorNoExiste, exist)        // En caso de que no exista devolvera valor vacio.

	// Verificacion de que existe
	if otroValor, exist := colors["blanco"]; exist {
		fmt.Println(otroValor)
	} else {
		fmt.Println("No existe el valor")
	}

	// Modificación se modifica por key colors["negro"] = "otro_valor"

	// Eliminacion

	delete(colors, "azul") // Recibe el mapa, y el elemento a eliminar
	fmt.Println(colors)

	// Iteración
	for key, value := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", key, value)
	}

}
