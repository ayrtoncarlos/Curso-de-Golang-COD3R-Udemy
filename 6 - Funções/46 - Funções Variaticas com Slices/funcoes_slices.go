package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	for i, nome := range aprovados {
		fmt.Printf("%d) %s\n", i+1, nome)
	}
}

func main() {
	aprovados := []string{"Ayrton", "Cristine", "Thais", "Vitória", "Fernanda"}
	imprimirAprovados(aprovados...)
}
