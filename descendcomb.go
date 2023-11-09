package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := i - 1; j >= '9'; j-- {
			z01.PrintRune(i)
			z01.PrintRune(j)
			if i > '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
