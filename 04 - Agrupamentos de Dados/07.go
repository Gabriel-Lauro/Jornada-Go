/*
Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
package main

import "fmt"

func main() {
	slice := [][]string{
		[]string{"Gabriel", "Lauro", "Grêmio"},
		[]string{"Zézin", "Silva", "Soltar Pipa"},
		[]string{"Matvey", "Safonov", "Pegar Penalti"},
	}
	for i, _ := range slice {
		for j, _ := range slice {
			fmt.Print(slice[i][j], "\t")
		}
		fmt.Println("")
	}
}
