package animal

import "fmt"

/* 	Para cumplir con la interfaz lo unico que tienen que tener las clases es alguna implementación
de los metodos del contrato de la interfaz
*/

type Animal interface {
	Sound()
}

func MakeSound(a Animal) {
	a.Sound()
}

type Dog struct {
	Name string
}

func NewDog(name string) *Dog {
	return &Dog{Name: name}
}

func (a Dog) Sound() {
	fmt.Printf("%s say Guau guau\n", a.Name)
}

type Cat struct {
	Name string
}

func NewCat(name string) *Cat {
	return &Cat{Name: name}
}

func (a Cat) Sound() {
	fmt.Printf("%s say miau miau\n", a.Name)
}

type Bird struct {
	Name string
}

func NewBird(name string) *Bird {
	return &Bird{Name: name}
}

func (a Bird) Sound() {
	fmt.Printf("%s say pio pio\n", a.Name)
}
