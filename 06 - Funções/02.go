/*
Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
Passe um valor do tipo slice of int como argumento para a função;
Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
Passe um valor do tipo slice of int como argumento para a função.
*/
package main

import "fmt"

func main() {
	fmt.Println(somaNumeros(1, 2, 3, 8))
	fmt.Println(somaSlice([]int{1, 2, 3, 8}))
}

func somaNumeros(x ...int) int {
	total := 0
	for i, _ := range x {
		total += x[i]
		i++
	}
	return total
}

func somaSlice(x []int) int {
	total := 0
	for _, num := range x {
		total += num
	}
	return total
}
