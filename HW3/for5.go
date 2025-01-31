package main

import "fmt"

func main() {
	for i := 1; i <= 8; i++ {
		if i%2 != 0 {
			fmt.Printf("#")
		} else {
			fmt.Printf(" ")
		}
	}
}
