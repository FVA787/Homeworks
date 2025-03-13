package main

import (
	package2 "Functions/pointer2"
	"fmt"
)

func main() {
	a := 2

	fmt.Println(a)

	fmt.Println(
		package2.Sum(&a),
	)

	fmt.Println(a)
}
