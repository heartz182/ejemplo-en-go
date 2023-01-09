//Operadores lógicos y de comparación

package main

import "fmt"

func main() {
	n := 2 == 3
	m := 5 == 5
	fmt.Printf("%v, %T", n, n)
	fmt.Printf("%v, %T\n", m, m)

	//Representación de operadores binarios, lógicosmy aritméticos

	a := 5 //0101
	b := 3 //0011
	//0 is false, 1 is true
	fmt.Println(a & b)  // Se unifican los dos valores con 'AND', 0001
	fmt.Println(a | b)  //se unifican los dos valores con 'OR', 0111
	fmt.Println(a ^ b)  // Se unifican los dos valores con 'XOR' (bicondicional), 0110
	fmt.Println(a &^ b) //se unifican los dos valores con 'and/or', 0100 , entender mejor
	y := 5
	fmt.Println((y << 4)) // en representación de exponenciales (binarios), se añade 4 al exponente
	fmt.Println(y >> 2)   //en representación de exponenciales (binarios), se resta 2 al exponentego
}
