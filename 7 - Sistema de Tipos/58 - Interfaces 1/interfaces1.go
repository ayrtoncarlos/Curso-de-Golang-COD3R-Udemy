package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$%.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Ayrton", "Andrade"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Lápis", 1.35}
	imprimir(p2)
}
