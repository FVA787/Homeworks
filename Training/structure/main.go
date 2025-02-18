package main

import (
	"fmt"
	"structure/struc"
)

func main() {
	type User struct {
		Lastname string
	}
	user := struc.Vova{
		Name: "Alice",
		Age:  25,
	}
	var user3 User
	user3.Lastname = "Alice"

	var a rune = 'j'

	var x int
	x = 55
	var ptr *int
	ptr = &x

	s := 5 // 101 в двоичном виде
	g := 3 // 011 в двоичном виде

	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(len(fruits))
	fmt.Println(fruits[len(fruits)-1])
	fmt.Println()
	for i := len(fruits) - 1; i >= 0; i-- {
		fmt.Println(fruits[i]) // "orange", "banana", "apple"
	}

	fmt.Println(user)
	fmt.Println(user3)
	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(10 % 4)
	fmt.Println(s | g)

}
