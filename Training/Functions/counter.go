package main

import "fmt"

func counterasd() (func() int, error) { // здесь функция counterasd() возвращает другую функцию func() int {}
	i := 0              // переменная, замкнутая в области видимости функции func() int {....}
	return func() int { //Здесь мы не вызываем функцию (чтобы вызвать функцию нужно написать круглые скобочки (), а содаем ее, но одновременно пишем команду что мы ретерним. команда ретерн и создание функции одновременно, но ретерн и создание функции это разные операции, они не имет нтношения друг к другу
		i++
		return i
	}, nil
}

//func counterasd() func() int { // здесь функция counterasd() возвращает другую функцию func() int {}
//	i := 0 // переменная, замкнутая в области видимости функции func() int {....}
//	myfunc_which_I_want_to_return := func() int {
//		i++
//		return i
//	}
//	return myfunc_which_I_want_to_return
//}

func main() {
	increment, _ := counterasd() // создаем замыкание   "_" - это игнорирование возвращаемого значения. Типа мы говорим что мы знаем что она возвращает два значения, но нам не важно второе, это пишется чтобы соблюсти контракт (это штука которую обязуются выполнять 2 стороны - каунтерасд обязуется вернуть 2 значекния, принммающая сторона обязуется принять тоже 2 знаенчения
	fmt.Println(increment)
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println()
	newCounter, _ := counterasd() // новый счетчик, независимый от первого
	fmt.Println(newCounter())
	fmt.Println(newCounter())
	fmt.Println(newCounter())
}
