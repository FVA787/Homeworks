package main

import "fmt"

func main() {
	var tables = 8
	var lines = 8
	for i := 1; i <= lines; i++ {
		if i%2 != 0 {
			for i := 1; i <= tables; i++ {
				if i%2 != 0 {
					if i == tables {
						fmt.Printf("#\n")
					} else {
						fmt.Printf("#")
					}
				} else {
					if i == tables {
						fmt.Printf(" \n")
					} else {
						fmt.Printf(" ")
					}
				}
			}
		} else {
			for i := 1; i <= tables; i++ {
				if i%2 != 0 {
					if i == tables {
						fmt.Printf(" \n")
					} else {
						fmt.Printf(" ")
					}
				} else {
					if i == tables {
						fmt.Printf("#\n")
					} else {
						fmt.Printf("#")
					}
				}
			}
		}
	}
}
