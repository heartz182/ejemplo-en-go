package main

import (
	"fmt"
	"math"
)

func main() {

	cilindro()
	esfera()
	cono()
}

func cilindro() {
	var r, h float64

	pi := math.Pi

	fmt.Println("Ingrese el radio (r)")
	fmt.Scanf("%f", &r)

	fmt.Println("Ingrese la altura (h)")
	fmt.Scanf("%f", &h)

	vol := (pi * (r * r)) * h

	fmt.Println("El volumen del cilindro es:", vol)

}

func esfera() {
	var r float64

	pi := math.Pi

	fmt.Println("Ingrese el radio (r)")
	fmt.Scanf("%f", &r)

	vol := ((4 * pi) * (r * r * r)) / 3

	fmt.Println("El volumen de la esfera es:", vol)
}

func cono() {
	var r, h float64

	pi := math.Pi

	fmt.Println("Ingrese el radio (r)")
	fmt.Scanf("%f", &r)

	fmt.Println("Ingrese la altura (h)")
	fmt.Scanf("%f", &h)

	vol := (pi * (r * r) * h) / 3

	fmt.Println("El volumen del cono es:", vol)
}
