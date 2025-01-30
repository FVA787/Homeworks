package main

import (
	"fmt"
	"math/rand/v2"
)

func twoRandom() (int, int) {
	return rand.IntN(2), rand.IntN(2)
}
func main() {
	x, y := twoRandom()
	fmt.Println("IN MAIN: ", x, y)
	if x, y := twoRandom(); x < y {
		fmt.Println("IN IF: ", x, y)
		fmt.Println("x equals ", x, "y equals ", y)
	} else {
		fmt.Println("IN IF: x >= y")
		fmt.Println("x equals ", x, "y equals ", y)
	}
	fmt.Println("IN MAIN PRINT ONE MORE TIME: ", x, y)
}
