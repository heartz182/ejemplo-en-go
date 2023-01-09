package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese la operación (suma, de la forma numero + numero, ej: 2+2):")
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println("La operación ingresada es", operacion)
	valores := strings.Split(operacion, "+")
	fmt.Println("Estos son los valores ingresados", valores)
	fmt.Println("Primer y segundo valor sumados como texto:", valores[0]+valores[1])
	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])

	fmt.Println("Suma de los dos operadores matematicamente: ", operador1+operador2)

}

/*
Lo primero que se va hacer es un escáner para recibir los inputs del usuario. Para esto usaremos
una librería llamada bufio y su método NewScanner() que genera una nueva instancia de escaneo. Este
escaneo puede ser desde diferentes inputs de entrada, en este caso le pasaremos como parámetro os.Stdin para
que funcione con los inputs de la terminal.

Para que la terminal funcione necesitamos un par más de funciones:
scanner.Scan() recibe y guarda en la instancia del escáner el texto escrito en la terminal hasta un Enter.
scanner.Text() extrae el texto del escáner y lo podemos asignar a alguna variable.

El siguiente paso es extraer los valores y la operación que se está realizando. Para esto usamos la función
split perteneciente al paquete strings. Lo que hace es partir un string por medio de un caracter definido
como el “separador”.

La función strings.Split() regresa un array, por lo que una vez en nuestras variables podemos acceder a sus
 valores por medio de sus índices. De igual forma, la función únicamente hace la partición del string original,
 mas no su valor numérico. Hay que hacer un casting o transformación utilizando la función Atoi perteneciente al
 paquete strconv que convierte un valor string a entero.

Otra peculiaridad de Go es que las funciones pueden regresar más de un solo valor, como es el caso de strconv.Atoi().
Esta función a su vez regresa un error si algo sale mal, sin embargo, podemos utilizar guión bajo _ para omitir el uso
de la variable.
*/
