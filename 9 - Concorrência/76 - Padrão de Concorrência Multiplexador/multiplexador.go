package main

import (
	"fmt"

	"github.com/cod3rcursos/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar: Misturar (mensagens) em um canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.urionlinejudge.com.br/judge/pt", "https://codeforces.com/"),
		html.Titulo("https://www.geeksforgeeks.org/", "https://cp-algorithms.com/"),
	)

	fmt.Printf("%s | %s\n", <-c, <-c)
	fmt.Printf("%s | %s\n", <-c, <-c)
}
