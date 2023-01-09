package main

import "fmt"

func main() {

	var i, a int

	fmt.Println("¿Cual es tu edad?")
	fmt.Scanf("%d", &i)

	fmt.Println("¿en que año estas?")
	fmt.Scanf("%d", &a)

	n := a - i

	fmt.Println("naciste en el año", n)

}
