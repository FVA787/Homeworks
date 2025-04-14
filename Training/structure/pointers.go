package main

import "fmt"

func main() {
	x := 2524  // тип int
	xPtr := &x // тип *int
	var p *int // тип *int, значение nil

	ddx := new(int)
	*ddx = 42

	i := uint8(0)
	i--
	fmt.Println(i) // 255
	i -= 5
	fmt.Println(i) // 250

	fmt.Println(xPtr)  // 0xc000090020
	fmt.Println(*xPtr) // 2524
	fmt.Println(p)     // nil
	fmt.Println(*ddx)  // 42
	asd()              // "asd asd asd"
}

func asd() {
	fmt.Println("asd asd asd")
}
