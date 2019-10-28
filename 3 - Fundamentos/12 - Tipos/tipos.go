package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(10000))

	// Sem sinal (Só positivos)... uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// Com sinal... int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	// Representa um mapeamento da tabela unicode
	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor do i2 é", i2)

	// Números reais (float32, float64)
	// var x = float32(49.99)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá mundo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho de s1 é", len(s1))

	// String com múltiplas linhas
	s2 := `Olá
	mundo
	!
	`
	fmt.Println("O tamanho de s2 é", len(s2))
}
