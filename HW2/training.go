package main

import (
	"fmt"
	"math/rand/v2"
)

func isZero(i int) bool {
	if i == 0 {
		return true
	}
	return false
}
func main() {
	randNuRRmber := rand.IntN(5)
	fmt.Println(randNuRRmber)
	zero := isZero(randNuRRmber)
	if zero == true {
		fmt.Println("Зашли в IF")
	}
}
