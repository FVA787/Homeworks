package main

import "fmt"

func multiply(x int, y int) int { //функция как значение
	return x * y
}
func main() {
	operation := multiply // вводим переменную, которая будет равняться функции
	result := operation(2, 3)
	fmt.Println(result)
}
