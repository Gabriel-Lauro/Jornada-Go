/*
Crie um map com key tipo string e value tipo []string.
Key deve conter nomes no formato sobrenome_nome
Value deve conter os hobbies favoritos da pessoa
Demonstre todos esses valores e seus índices.
*/
package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"Gabriel_Lauro": []string{"Grêmio", "Secar o inter"},
		"Zezinho_Silva": []string{"hobbie1", "hobbie2", "hobbie3"},
	}

	pessoas["Neymar_JR"] = []string{"hobbie1", "hobbie2", "hobbie3"}

	for k, v := range pessoas {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}
