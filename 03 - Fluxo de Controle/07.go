/*
Utilizando a solução anterior, adicione as opções else if e else.
*/
package main

import "fmt"

func main() {
	x := 60
	if x < 20 {
		fmt.Println("Menor que 20")
	} else if x == 30 {
		fmt.Println("X = a 30")
	} else {
		fmt.Println("X é maior que 20 e diferente de 30")
	}

}
