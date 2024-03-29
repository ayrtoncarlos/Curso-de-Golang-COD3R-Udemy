package main

import (
	"fmt"
	"time"

	"github.com/cod3rcursos/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// Estrutura de controle específica para concorrência.
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos Perderam!"
		//default:
		//	return "Sem resposta!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://codeforces.com/",
		"https://www.geeksforgeeks.org/",
		"https://cp-algorithms.com/",
	)

	fmt.Println(campeao)
}
