package main

import "fmt"

func main() {
	var alumnos [4]string

	fmt.Println("Cual es el nombre del primer alumno?")
	fmt.Scanf("%s", &alumnos[0])

	fmt.Println("Cual es el nombre del segundo alumno?")
	fmt.Scanf("%s", &alumnos[1])

	fmt.Println("Cual es el nombre del tercer alumno?")
	fmt.Scanf("%s", &alumnos[2])

	fmt.Println("Cual es el nombre del cuarto alumno?")
	fmt.Scanf("%s", &alumnos[3])

	fmt.Println("El primer alumno es", alumnos[0])
	fmt.Println("El segundo alumno es", alumnos[1])
	fmt.Println("El tercer alumno es", alumnos[2])
	fmt.Println("El cuarto alumno es", alumnos[3])

}
