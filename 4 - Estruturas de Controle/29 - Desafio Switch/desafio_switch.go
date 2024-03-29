package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 6 && n < 8:
		return "C"
	case n >= 3 && n < 6:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Nota Inválida!"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.6))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(2.3))
	fmt.Println(notaParaConceito(10.1))
	fmt.Println(notaParaConceito(-0.1))
}
