/*
Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Demonstre estes valores.
*/
package main

import "fmt"

const (
	a = iota + 2026
	b
	c
	d
	e
)

func main() {

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e)

}
