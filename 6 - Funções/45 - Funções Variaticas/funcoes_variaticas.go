package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	if len(numeros) > 0 {
		return total / float64(len(numeros))
	}
	return 0.0
}

func main() {
	fmt.Printf("Média: %.2f\n", media(7.8, 9.4, 6.8, 5.4, 8.3))
	fmt.Printf("Média: %.2f\n", media())
}
