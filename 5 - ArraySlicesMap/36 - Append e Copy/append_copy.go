package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(array, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
