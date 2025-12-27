/*
Callback: passe uma função como argumento a outra função.
*/

package main

import "fmt"

func main() {

	adição(multiplicação)

}

func multiplicação(x int) int {
	return x * 2
}

func adição(f func(int) int) {
	fmt.Println(f(20) + 10)
}
