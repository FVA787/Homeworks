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

	var v any
	v = 1003087
	q, ok := v.(int) // здесь мы достали из интерфейса (то же самое, что тип any) интовое значение, то есть привели значение в интерфейсе к типу int.
	if ok {
		fmt.Println("q is an integer:", q)
	}
	fmt.Println(v)
	fmt.Println(q)

	switch a.(type) {
	case int:
		fmt.Println("int is an integer, value is:", a)
	case string:
		fmt.Println("string is a string, value is:", a)
	case float64:
		fmt.Println("float64 is a float, value is:", a)
	case struc.Vova:
		fmt.Println("Type is Vova, value is", a)
	default:
		fmt.Println("Type is unknown")
	}

}
