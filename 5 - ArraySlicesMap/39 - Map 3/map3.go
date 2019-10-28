package main

import "fmt"

func main() {
	funcsESalarios := map[string]map[string]float64{
		"A": {
			"Ayrton": 2568.7,
			"Ana":    3567.3,
		},
		"C": {
			"Cristine": 4675.9,
		},
		"M": {
			"Mateus":  3456.2,
			"Matheus": 2345.6,
		},
		"T": {
			"Thais": 3546.5,
		},
	}

	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "A")

	for letra, funcs := range funcsESalarios {
		fmt.Println(letra, funcs)
	}

	for letra, funcs := range funcsESalarios {
		fmt.Printf("Letra: %s\n", letra)
		for nome, salario := range funcs {
			fmt.Printf("\tNome: %s | Salario: %.2f\n", nome, salario)
		}
	}
}
