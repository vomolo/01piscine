package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for _, str := range arguments {
		for _, th := range str {
			z01.PrintRune(th)
		}
		z01.PrintRune('\n')
	}
}
