package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (Interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Ayrton", "Olá, tudo bem?", 5)
	// fale("Cristine", "Tudo sim!", 5)

	// go fale("Ayrton", "Ei...", 500)
	// go fale("Thais", "Oi...", 500)
	// fmt.Println("Fim!")

	go fale("Ayrton", "Ei...", 10)
	fale("Cristine", "Oi...", 5)
}
