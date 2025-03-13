package main

import "fmt"

type Lot struct {
	Buy_price  float64
	Sale_price float64
}

func (d Lot) Profit() float64 {
	return d.Sale_price*0.93 - d.Buy_price
}
func main() {
	Novorossiysk := Lot{3825156, 8000000}                                // Вводим переменную лота - какие именно у него цена покупки и продажи
	fmt.Printf("Прибыль по проекту: %.1f рублей", Novorossiysk.Profit()) // Печать: Внутри синтаксиса:  Объект.Метод: Объект - Здание (Novorossiysk), метод - Profit,
	// который описан выше в функции func (d Lot) Profit() int  (return d.Sale_price - d.Buy_price). Метод - это функция внутри
	// структуры, которая может быть вызвана только структурой, и ниоткуда больше.
}
