package main

import (
	"fmt"
	"structure/struc"
)

func main() {
	var a interface{} // то же самое, что var a any. Any - это allias for interface {} Синтаксис выгдядит так: type имяИнтерфейса interface {названиеМетода(списокПараметров) списокТиповВозвращаемыхЗначений}

	a = 123
	fmt.Println(a)

	a = "Hello world"
	fmt.Println(a)

	a = struc.Vova{
		Name: "Alice",
		Age:  25,
	}

	fmt.Println(a)

	var b any
	b = "Privet kak dela"
	fmt.Println(b)

	b = "Hello world"
	fmt.Println(b)
	fmt.Println()
	//var a any
	//a = 123

}
