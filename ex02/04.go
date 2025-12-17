/*
Crie um programa que:
Atribua um valor int a uma variável
Demonstre este valor em decimal, binário e hexadecimal
Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
Demonstre esta outra variável em decimal, binário e hexadecimal
*/
package main

import "fmt"

func main() {
	x := int(200)
	fmt.Printf("Decimal:     %d\n", x)
	fmt.Printf("Binario:     %b\n", x)
	fmt.Printf("hexadecimal: %x\n\n", x)

	y := x << 1

	fmt.Printf("Decimal:     %d\n", y)
	fmt.Printf("Binario:     %b\n", y)
	fmt.Printf("hexadecimal: %x\n", y)
}
