package main

import "fmt"

func main() {
	a := 1
	b := func(a int) {
		a++
		fmt.Println(a)
	}
	fmt.Println(a)
	b(a)
	b(10)
	fmt.Println(a)
}
