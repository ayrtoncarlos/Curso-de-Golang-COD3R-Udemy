package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 300
	f.turboLigado = true

	fmt.Printf("O ferrari %s está com turbo ligado? %t", f.nome, f.turboLigado)
}
