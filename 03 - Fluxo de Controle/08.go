/*
Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/
package main

import "fmt"

func main() {
	x := -30

	switch {
	case x < 20:
		fmt.Println("Menor que 20")

	case x == 30:
		fmt.Println("X = a 30")

	default:
		fmt.Println("X é maior que 20 e diferente de 30")
	}

}
