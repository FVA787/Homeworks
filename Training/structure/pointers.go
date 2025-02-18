package main

import "fmt"

func main() {
	x := 2524  // тип int
	xPtr := &x // тип *int
	var p *int // тип *int, значение nil

	fmt.Println(xPtr)
	fmt.Println(*xPtr)
	fmt.Println(p)
}
