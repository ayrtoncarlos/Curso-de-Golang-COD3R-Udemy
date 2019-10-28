package main

import "fmt"

func obterResultado(n float64) string {
	nota := int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 7, 8:
		return "B"
	case 5, 6:
		return "C"
	case 3, 4:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida!"
	}
}

func main() {
	fmt.Println(obterResultado(9.5))
	fmt.Println(obterResultado(6.4))
	fmt.Println(obterResultado(2.6))
	fmt.Println(obterResultado(11.1))
	fmt.Println(obterResultado(-1.1))
}
