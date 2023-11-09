package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for i := 99; i >= 10; i-- {
		for j := i - 1; j >= 10; j-- {
			z01.PrintRune(rune(i/10 + '0'))
			z01.PrintRune(rune(i%10 + '0'))
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10 + '0'))
			z01.PrintRune(rune(j%10 + '0'))

			if j != 10 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
