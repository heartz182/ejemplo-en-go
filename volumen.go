package main

import (
	"fmt"
	"math"
)

func main() {
	var r, h float64

	pi := math.Pi

	fmt.Println("Ingrese el radio (r)")
	fmt.Scanf("%f", &r)

	fmt.Println("Ingrese la altura (h)")
	fmt.Scanf("%f", &h)

	vol := (pi * (r * r)) * h

	fmt.Println("El volumen es:", vol)

}
