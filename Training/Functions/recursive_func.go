package main

import "fmt"

func vova(n int) int {
	if n == 0 {
		return 0
	}
	return n + vova(n-1)
}

func main() {
	fmt.Println(vova(11))
}
