package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * float32(math.Pi)
}

func main() {
	fmt.Println(Square{4})
	fmt.Println(Circle{4})
}
