package main

import "fmt"

func main() {
	var a, b, c, d, e, mayor int

	fmt.Println("Ingrese el primer numero entero")
	fmt.Scanf("%d", &a)

	fmt.Println("Ingrese el segundo valor entero")
	fmt.Scanf("%d", &b)

	fmt.Printf("Ingrese el tercer valor entero")
	fmt.Scanf("%d", &c)

	fmt.Scanf("Ingrese el cuarto valor entero")
	fmt.Scanf("%d", &d)

	fmt.Println("Ingrese el quinto valor entero")
	fmt.Scanf("%d", &e)

	if a > b {
		mayor = a
	} else {
		mayor = b
	}
	if c > mayor {
		mayor = c
	}
	if d > mayor {
		mayor = d
	}
	if e > mayor {
		mayor = e

	}
	fmt.Println("el numero mayor es:", mayor)

}
