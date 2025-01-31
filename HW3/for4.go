package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c"}
	for a, letter := range letters {
		fmt.Printf("Index: %d Value: %s; Vova %s\n", a, letter, "Molodets")
	}
}
