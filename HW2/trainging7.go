package main

import "fmt"

func main() {
	var s interface{}
	s = "# # #"
	switch stype := s.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("%T", stype)
	}
}
