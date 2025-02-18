package main

import "fmt"

func main() {
	var a interface{} // то же самое, что var a any. Any - это allias for interface {}

	a = 123
	fmt.Println(a)

	a = "Hello world"

	fmt.Println(a)

	var b any
	b = 123
	fmt.Println(b)

	b = "Hello world"
	fmt.Println(b)
}
