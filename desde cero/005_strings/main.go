package main

import "fmt"

func main() {

	var nombre, apellidos string
	var peso float32

	fmt.Println("Cual es tu nombre?")
	fmt.Scanf("%s", &nombre)

	fmt.Println("Cual es tu apellido")
	fmt.Scanf("%s", &apellidos)

	fmt.Println("Cual es tu peso?")
	fmt.Scanf("%f", &peso)

	fmt.Println("El nombre es: ", nombre, apellidos)
	fmt.Println("El peso es: ", peso, "kg")

}
