package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b //возвращаем результат сложения
}

func main() {
	result := add(3, 4) // Вызов функции add с аргументами 3 и 4
	fmt.Println(result)
}
