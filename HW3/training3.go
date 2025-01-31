package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randNumber := rand.IntN(5)
	if randNumber == 1 {
		fmt.Println(randNumber)
		fmt.Println("Первый")
	} else if randNumber == 2 {
		fmt.Println(randNumber)
		fmt.Println("Второй")
	} else if randNumber == 3 {
		fmt.Println(randNumber)
		fmt.Println("Третий")
	} else {
		fmt.Println(randNumber)
		fmt.Println("Какой-то")
	}
}
