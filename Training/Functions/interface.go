package main

import (
	"fmt"
	"math"
)

type shape interface { //объявляем интерфейс, который будет иметь только 1 метод area, который возвращает площадь фигуры с типом float32
	area() float32
}

type square struct { // объявляем тип square - структуру, которая имеет тоолько одно поле sideLength с типом float32
	sideLength float32
}

func (s square) area() float32 { // функция с аргументом в виде переменной s (мы ее не объявляли, но здесь же, видимо, мы ее и объявляем),
	// которая (переменная) имеет тип square
	return s.sideLength * s.sideLength // поскольку у типа square есть единственный параметр sideLength мы пишем s.sideLength
}

type circle struct { // эти две структуры по умолчанию реализуют интерфейс shape, поскольку мы объявиил для них метод area с возвращаемым значением типа float32
	radius float32
}

func (c circle) area() float32 {
	return c.radius * c.radius * float32(math.Pi)
}

func main() {
	square := square{sideLength: 5}
	circle := circle{radius: 5}

	fmt.Println(square.area())
	fmt.Println(circle.area())
}

func printShapeArea(s shape) {

}
