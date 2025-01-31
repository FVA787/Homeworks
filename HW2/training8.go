package main

import "fmt"

func main() {
	switch age, _, _ := exampleIf.FindUserBySNILS("000-000-000 00"); age {
	case 30, 25:
		fmt.Println(30)
		fmt.Println(25)
	case 14:
		fmt.Println(14)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Age = %d\n", age)
	}
}
