package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) obterValorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	var pedido1 = pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 10, preco: 4.75},
			item{2, 5, 6.60},
			item{6, 100, 0.75},
		},
	}
	fmt.Printf("O valor total é R$%.2f\n", pedido1.obterValorTotal())
}
