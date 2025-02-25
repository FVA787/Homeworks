package main

import "fmt"

type Rctg struct {
	width  int
	height int
}

func (d Rctg) Area123() int {
	return d.width * d.height
}
func main() {
	rect := Rctg{10, 5}         // Вводим переменную ректангла - какие именно у него длина и ширина
	fmt.Println(rect.Area123()) // Печать: Внутри синтаксиса:  Объект.Метод Объект - ректангл, метод - Area123,
	// который описан выше в функции func (d Rctg) Area123() int  (return d.width * d.height)
}
