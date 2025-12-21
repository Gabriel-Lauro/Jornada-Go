/*
Usando uma literal composta:
Crie uma slice de tipo int
Atribua 10 valores a ela
Utilize range para demonstrar todos estes valores.
E utilize format printing para demonstrar seu tipo.
*/
package main

import "fmt"

func main() {
	slice := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for i, _ := range slice {
		fmt.Print(slice[i], " ")
	}

	fmt.Printf("\n\nO Slice %v é do tipo : %T", slice, slice)

}
