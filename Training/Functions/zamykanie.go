package main

import "fmt"

func main() {
	makeFn := func() func(int) int {
		count := 0
		return func(i int) int {
			count++
			return count * i
		}
	}

	fn1, fn2 := makeFn(), makeFn()

	fmt.Printf("fn1 : ")
	for i := 1; i <= 4; i++ {
		fmt.Printf("%d, ", fn1(i))
	}
	fmt.Println()
	fmt.Printf("fn2 : ")
	a := -4
	for i := -1; i >= a; i-- {
		if i > a {
			fmt.Printf("%d, ", fn2(i))
		} else {
			fmt.Printf("%d.", fn2(i))
		}
	}
}
