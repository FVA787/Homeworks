package main

import "fmt"

func counter() func() int {
	count := 0 // переменная, замкнутая в области видимости
	return func() int {
		count++
		return count
	}
}

func main() {
	increment := counter() // создаем замыкание
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	newCounter := counter() // новый счетчик, независимый от первого
	fmt.Println(newCounter())
	fmt.Println(newCounter())
	fmt.Println(newCounter())
}
