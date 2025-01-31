package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("i = %d\n", i)
		switch i {
		case 2:
			break
		case 3:
			fmt.Println("It's three")
		default:
			fmt.Println("Ok")
		}
		fmt.Println(".")
	}
}
