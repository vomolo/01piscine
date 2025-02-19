package main

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	var result int = 1
	if nb > 63 || nb < -63 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	for i := -1; i >= nb; i-- {
		result = result * i
	}
	return result
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
