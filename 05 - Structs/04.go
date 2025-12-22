/*
Crie e use um struct anônimo.
Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
package main

import "fmt"

func main() {

	x := struct {
		mapaX  map[string]string
		sliceX []int
	}{
		mapaX: map[string]string{
			"dasdas": "dsakdas",
		},
		sliceX: []int{1, 1, 1},
	}

	fmt.Print(x)

}
