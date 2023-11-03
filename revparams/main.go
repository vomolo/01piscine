package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	argsIndex := len(args) - 1
	for argsIndex >= 0 {
		for _, char := range args[argsIndex] {
			z01.PrintRune(char)
		}
		argsIndex--
		z01.PrintRune('\n')
	}
}
