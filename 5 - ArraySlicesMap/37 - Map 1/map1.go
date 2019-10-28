package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[987654321] = "Ayrton"
	aprovados[123456789] = "Cristine"
	aprovados[123459876] = "Thais"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d\n", nome, cpf)
	}

	fmt.Println(aprovados[123456789])
	delete(aprovados, 987654321)
	fmt.Println(aprovados)
}
