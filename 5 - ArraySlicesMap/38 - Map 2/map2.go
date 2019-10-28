package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Ayrton":   2500.0,
		"Cristine": 4567.6,
		"Thais":    3456.2,
	}

	fmt.Println(funcsESalarios)

	funcsESalarios["Mateus"] = 2756.8

	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
