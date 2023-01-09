package main

import "fmt"

// hola devuelve "Hola" seguido de la cadena de entrada s
func hola(s string) string {
	return "Hola " + s
}

func main() {
	// Imprime "Hola Eduardo" en la consola
	fmt.Println(hola("Eduardo"))
}
