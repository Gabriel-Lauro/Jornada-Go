/*
Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
Nome
Sobrenome
Sabores favoritos de sorvete
Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/
package main

import "fmt"

type pessoa struct {
	nome         string
	sobrenome    string
	saborSorvete []string
}

func main() {
	pessoa1 := pessoa{
		nome:         "Zezinho",
		sobrenome:    "silva",
		saborSorvete: []string{"morango", "flocos"},
	}
	pessoa2 := pessoa{
		nome:         "Gabriel",
		sobrenome:    "Lauro",
		saborSorvete: []string{"Azul", "Morango"},
	}
	fmt.Println(pessoa1, pessoa2)

	for _, v := range pessoa1.saborSorvete {
		fmt.Printf("%v - ", v)
	}

	for _, v := range pessoa2.saborSorvete {
		fmt.Printf("%v - ", v)
	}
}
