package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case bool:
		return "Booleano"
	case func():
		return "Função"
	default:
		return "Tipo desconhecido"
	}
}

func main() {
	fmt.Println(tipo(2.6))
	fmt.Println(tipo(5))
	fmt.Println(tipo("Ayrton"))
	fmt.Println(tipo(true))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
