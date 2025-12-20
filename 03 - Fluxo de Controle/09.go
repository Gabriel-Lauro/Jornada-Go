/*
Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/

package main

import "fmt"

func main() {
	esporteFavorito := "volei"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("O esporte vaforito é futebol")

	case "volei":
		fmt.Println("O esporte vaforito é volei")

	default:
		fmt.Println("Esporte favorito desconhecido")
	}

}