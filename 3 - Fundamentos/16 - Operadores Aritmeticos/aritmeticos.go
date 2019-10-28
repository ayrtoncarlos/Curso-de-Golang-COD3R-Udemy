package main

import (
	"fmt"
	"math"
)

func main() {
	a := 5
	b := 8

	fmt.Println("Soma -->", a+b)
	fmt.Println("Substração -->", a-b)
	fmt.Println("Multiplicação -->", a*b)
	fmt.Println("Divisão -->", a/b)
	fmt.Println("Resto -->", a%b)

	// Bitwise
	fmt.Println("E -->", a&b)
	fmt.Println("Ou -->", a|b)
	fmt.Println("Xor -->", a^b)

	// Outras operações utilizando math..
	c := 2.0
	d := 4.0

	fmt.Println("Maior -->", math.Max(float64(a), float64(b)))
	fmt.Println("Menor -->", math.Min(c, d))
	fmt.Println("Exponenciação -->", math.Pow(c, d))
}
