/*
Crie um loop utilizando a sintaxe: for condition {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/
package main

import "fmt"

func main() {

	ano1 := 2002
	ano2 := 2025

	for {
		fmt.Println(ano1)
		if ano1 == ano2 {
			break
		}
		ano1++
	}

}
