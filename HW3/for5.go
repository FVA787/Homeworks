package main

import "fmt"

func main() {
	for i := 1; i <= 8; i++ {
		if i%2 != 0 {
			for i := 1; i <= 8; i++ {
				if i%2 != 0 {
					if i == 8 {
						fmt.Printf("#\n")
					} else {
						fmt.Printf("#")
					}
				} else {
					if i == 8 {
						fmt.Printf(" \n")
					} else {
						fmt.Printf(" ")
					}
				}
			}
		} else {
			for i := 1; i <= 8; i++ {
				if i%2 != 0 {
					if i == 8 {
						fmt.Printf(" \n")
					} else {
						fmt.Printf(" ")
					}
				} else {
					if i == 8 {
						fmt.Printf("#\n")
					} else {
						fmt.Printf("#")
					}
				}
			}
		}
	}
}
