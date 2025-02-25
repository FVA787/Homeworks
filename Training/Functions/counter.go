package main

import "fmt"

func counterasd() func() int { // здесь функция counterasd() возвращает другую функцию func() int {}
	i := 0 // переменная, замкнутая в области видимости функции func() int {....}
	return func() int {
		i++
		return i
	}
}

func main() {
	increment := counterasd() // создаем замыкание
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println()
	newCounter := counterasd() // новый счетчик, независимый от первого
	fmt.Println(newCounter())
	fmt.Println(newCounter())
	fmt.Println(newCounter())
}
