package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	rating := rand.IntN(5) + 1
	switch rating {
	case 1:
		fmt.Println("Кол, на пересдачу")
	case 2:
		fmt.Println("Двойка, на пересдачу")
	case 3:
		fmt.Println("Тройка, давай до свидания")
	case 4:
		fmt.Println("Все идеально, но ты мог лучше, четверка")
	case 5:
		fmt.Println("Отлично")
	default:
		fmt.Println("INCORRECT")
	}
}
