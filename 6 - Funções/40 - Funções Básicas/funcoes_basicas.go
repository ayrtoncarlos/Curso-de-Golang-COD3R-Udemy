package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(param1 string, param2 string) {
	fmt.Printf("F2: %s %s\n", param1, param2)
}

func f3() string {
	return "F3"
}

func f4(param1, param2 string) string {
	return fmt.Sprintf("F4: %s %s", param1, param2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")

	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println(r51, r52)
}
