/*
Crie uma função que retorne um int
Crie outra função que retorne um int e uma string
Chame as duas funções
Demonstre seus resultados
*/
package main

import "fmt"

func main() {
	fmt.Println(superFunçãoInt(5))
	fmt.Println(superFunçãoInteString(8, "abacate"))
}

func superFunçãoInt(x int) int {
	return x * 10
}

func superFunçãoInteString(x int, y string) (int, string) {
	return x * 5, y + " estragado"
}
