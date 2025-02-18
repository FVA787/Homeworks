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
	fmt.Println(i)

	fmt.Println(xPtr)  // 0xc000090020
	fmt.Println(*xPtr) // 2524
	fmt.Println(p)     // nil
	fmt.Println(*ddx)
	asd()
}

func asd() {
	fmt.Println("asd asd asd")
}
