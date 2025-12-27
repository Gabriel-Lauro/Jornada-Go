/*Faça um fatorial usando recursividade*/
package main

import "fmt"

func main() {
	fmt.Print(fatorial(5))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}
