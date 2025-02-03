package main

import "fmt"

func main() {

	var lines, tables int
	fmt.Print("Enter the number of lines and tables (Format: lines, tables): ")
	fmt.Scanf("%d, %d", &lines, &tables)
	for i := 0; i < lines; i++ {
		for j := 0; j < tables; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println() //Функция Println в языке программирования Go автоматически добавляет символ новой строки в конце вывода.
	}
}
