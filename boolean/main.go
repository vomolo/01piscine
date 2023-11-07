package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func even(numm int) int {
	if numm%2 == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	lengthOfArg := os.Args[1:]
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(len(lengthOfArg)) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
