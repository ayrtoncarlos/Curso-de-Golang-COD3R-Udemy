package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(3, 5))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(3, 5))
}
