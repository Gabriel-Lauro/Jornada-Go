/*
Crie constantes tipadas e não-tipadas.
Demonstre seus valores.
*/
package main

import "fmt"

const x float64 = 10
const y = 10

func main() {
	fmt.Printf("Tipada :    %v, %T\n", x, x)
	fmt.Printf("Não Tipada: %v, %T", y, y)
}
