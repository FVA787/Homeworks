package main

import (
	"fmt"
)

func main() {
	add := func(a int, b int) int { // Это функция без имени или лямбда-функция или анонимная функция (без имени), то есть мы не присваиваем
		// ей имя, мы просто пишем func (аргумент, тип аргумента) тип возвращаемого значения.
		return a * b
	}
	result := add(3, 4)
	fmt.Println(result)
}
