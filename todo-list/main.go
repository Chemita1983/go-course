// Ej: Gestor de tareas

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (lt *ListaTareas) agregarTarea(t Tarea) {
	lt.tareas = append(lt.tareas, t)
}

func (lt *ListaTareas) marcarTareaCompletada(name string) {
	if index, exist := lt.existeTarea(name); exist {
		lt.tareas[index].completado = true
		return
	}

	fmt.Println("La tarea introducida no existe")
}

func (lt *ListaTareas) eliminarTarea(name string) {
	if index, exist := lt.existeTarea(name); exist {
		lt.tareas = append(lt.tareas[:index], lt.tareas[index+1:]...)
		return
	} else {
		fmt.Println("La tarea introducida no existe")

	}
}

func (lt *ListaTareas) editarTarea(name string, tareaModificada Tarea) {
	if index, exist := lt.existeTarea(name); exist {
		if lt.tareas[index].completado == true {
			fmt.Println("La tarea no se puede modificar, ha sido completada")
			return
		}
		lt.tareas[index] = tareaModificada
		return
	}

	fmt.Println("La tarea introducida no existe")
}

func (lt *ListaTareas) existeTarea(name string) (index int, exist bool) {
	for index := range lt.tareas {
		if lt.tareas[index].nombre == name {
			return index, true
		}
	}
	return 0, false
}

func main() {

	listaTareas := ListaTareas{}

	// Instancia para la entrada de datos para leer datos debido a limitaciones
	leer := bufio.NewReader(os.Stdin)
	// Menu de navegacion
	for {
		var opcion int
		fmt.Println("")
		fmt.Println("Seleccione una opcion: \n",
			"1. Agregar tarea \n",
			"2. Marcar tarea como completada \n",
			"3. Editar tarea \n",
			"4. Eliminar tarea \n",
			"5. Salir")

		fmt.Scanln(&opcion)

		fmt.Println("")
		switch opcion {
		case 1:
			var nuevaTarea Tarea
			fmt.Print("Introduzca el nombre de la nueva tarea: ")
			nuevaTarea.nombre, _ = leer.ReadString('\n')
			fmt.Print("Introduzca la descripcion de la nueva tarea: ")
			nuevaTarea.desc, _ = leer.ReadString('\n')
			listaTareas.agregarTarea(nuevaTarea)
		case 2:
			var nombreTarea string
			fmt.Println("Introduzca el nombre de la tarea a completar: ")
			nombreTarea, _ = leer.ReadString('\n')
			listaTareas.marcarTareaCompletada(nombreTarea)
		case 3:
			var nombreTarea string
			fmt.Print("Introduzca el nombre de la tarea a editar: ")
			nombreTarea, _ = leer.ReadString('\n')
			var nuevaTarea Tarea
			fmt.Print("Introduzca el nuevo nombre de la  tarea: ")
			nuevaTarea.nombre, _ = leer.ReadString('\n')
			fmt.Print("Introduzca la nueva descripcion de la tarea: ")
			nuevaTarea.desc, _ = leer.ReadString('\n')
			listaTareas.editarTarea(nombreTarea, nuevaTarea)
		case 4:
			var nombreTarea string
			fmt.Print("Introduzca el nombre de la tarea a eliminar: ")
			nombreTarea, _ = leer.ReadString('\n')
			listaTareas.eliminarTarea(nombreTarea)
		case 5:
			fmt.Println("Hasta la próxima..")
			return
		default:
			fmt.Println("Opcion no valida")
		}

		fmt.Println("Lista de tareas")
		fmt.Println("====================================================")

		for index, value := range listaTareas.tareas {
			fmt.Printf("%d - %s - %s - Completado: %t\n", index, value.nombre, value.desc, value.completado)
		}

		fmt.Println("====================================================")
	}
}
