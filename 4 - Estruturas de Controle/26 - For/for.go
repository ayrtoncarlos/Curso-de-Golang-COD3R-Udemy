package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println("\nMisturando..")

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando..")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Par ")
		} else {
			fmt.Printf("Impar ")
		}
	}

	fmt.Println("\nMisturando..")

	for {
		fmt.Println("Para sempre!")
		time.Sleep(time.Second)
	}
}
