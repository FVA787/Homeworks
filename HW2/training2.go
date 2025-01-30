package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randNumber := rand.IntN(5)
	fmt.Println(randNumber)

	if randNumber == 1 {
		fmt.Println("Первый")
	} else {
		fmt.Println("Не первый")
	}

}
