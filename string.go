// Strings y bytes
package main

import "fmt"

func main() {
	s := "soy una cadena"
	t := "estoy junto a una cadena"
	fmt.Printf("%v\n", s+t)
	b := []byte(t)
	fmt.Printf("%v, %T", b, b)
}
