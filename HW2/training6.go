package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Second() < 15:

		fmt.Printf("<15 %v\n", t)
	case t.Second() < 30:
		fmt.Printf("<30 %v\n", t)
	case t.Second() < 45:
		fmt.Printf("<45 %v\n", t)
	default:
		fmt.Printf(">45 %v\n", t)
	}
}
