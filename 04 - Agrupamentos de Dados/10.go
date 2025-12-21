/*
Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
*/
package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"Gabriel_Lauro": []string{"Grêmio", "Secar o inter"},
		"Zezinho_Silva": []string{"hobbie1", "hobbie2", "hobbie3"},
	}

	pessoas["Neymar_JR"] = []string{"hobbie1", "hobbie2", "hobbie3"}

	delete(pessoas, "Zezinho_Silva")

	for k, v := range pessoas {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}
