/*
Crie um tipo "retangulo"
Crie um tipo "círculo"
Crie um método "área" para cada tipo que calcule e retorne a área da figura
  - Área do círculo: π * raio^2
  - Área do retângulo: largura * altura
Crie um método "perímetro" para cada tipo que calcule e retorne o perímetro da figura
  - Perímetro do círculo: 2 * π * raio
  - Perímetro do retângulo: 2 * (largura + altura)
Crie um tipo "figura" que defina como interface qualquer tipo que tiver os métodos "área" e "perímetro"
Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
Crie uma função "perimetroInfo" que receba um tipo "figura" e retorne o perímetro da figura
Crie um valor de tipo "retangulo"
Crie um valor de tipo "círculo"
Use a função "info" para demonstrar a área do "retangulo"
Use a função "info" para demonstrar a área do "círculo"
Use a função "perimetroInfo" para demonstrar o perímetro do "retangulo"
Use a função "perimetroInfo" para demonstrar o perímetro do "círculo"
*/

package main

import (
	"fmt"
	"math"
)

type figura interface {
	area() float64
	perimetro() float64
}

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

func (q retangulo) area() float64 {
	return q.altura * q.largura
}

func (q retangulo) perimetro() float64 {
	return (q.largura * 2) + (q.altura * 2)
}

func (c circulo) area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}

func (c circulo) perimetro() float64 {
	return c.raio * 2 * math.Pi
}

func info(f figura) {
	fmt.Printf("%T\nArea: %v\nPerimetro %v\n\n", f, f.area(), f.perimetro())
}

func main() {
	retangulo1 := retangulo{20, 10}
	circulo1 := circulo{25}

	info(circulo1)
	info(retangulo1)
}
