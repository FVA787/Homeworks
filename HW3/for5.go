package main

import "fmt"

func main() {
	for i := 1; i <= 8; i++ {
		for i := 1; i <= 8; i++ {
			if i%2 != 0 {
				if i%8 == 0 {
					fmt.Printf("#\n")
				} else {
					fmt.Printf("#")
				}
			} else {
				if i%8 == 0 {
					fmt.Printf(" \n")
				} else {
					fmt.Printf(" ")
				}
			}
		}
	}

}
