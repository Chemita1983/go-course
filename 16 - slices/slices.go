package main

import "fmt"

func main() {

	// Formas de inicializar slices
	var sl []int
	diasSemana := []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}
	diasEntreSemana := diasSemana[0:4] // Crear un slice a partir de otro

	fmt.Println(sl)
	fmt.Println(diasSemana)
	fmt.Println(diasEntreSemana)

	diasEntreSemana = append(diasEntreSemana, "Viernes")                            //Agregamos un nuevo elemento
	diasEntreSemana = append(diasEntreSemana, "Sabado", "Domingo", "Otro elemento") //Agregamos tantos elementos como queramos, aumenta su capacidad

	fmt.Println(diasEntreSemana)
	fmt.Println("Longitudes")
	fmt.Println(len(diasEntreSemana)) // Los elementos que tiene.
	fmt.Println(cap(diasEntreSemana)) // El total de elementos que tiene de capacidad

	diasEntreSemana = append(diasEntreSemana[:2], diasEntreSemana[2:5]...) // Podemos quitar slices
	fmt.Println(diasEntreSemana)

	// Otra manera de crear Slices
	nombres := make([]string, 5)
	nombres[0] = "Chema"
	nombres[1] = "Juan"
	fmt.Println(nombres)

	// Copiar slices
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("Elementos copiados", copy(slice2, slice1))

}
