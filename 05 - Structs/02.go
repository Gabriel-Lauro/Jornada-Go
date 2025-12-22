/*
Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
Demonstre os valores do map utilizando range.
Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/
package main

import "fmt"

type pessoa struct {
	nome         string
	sobrenome    string
	saborSorvete []string
}

func main() {

	mapa := make(map[string]pessoa)

	mapa["Lauro"] = pessoa{
		nome:         "Gabriel",
		sobrenome:    "Lauro",
		saborSorvete: []string{"Azul", "Morango"},
	}

	mapa["Silva"] = pessoa{
		nome:         "Zézinho",
		sobrenome:    "Silva",
		saborSorvete: []string{"Flocos", "Chocolate"},
	}

	for _, valor := range mapa {
		fmt.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.saborSorvete {
			fmt.Println("–", valor)
		}
		fmt.Println("")
	}

}
