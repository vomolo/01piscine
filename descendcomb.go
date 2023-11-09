package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '7'; i >= '0'; i-- {
		for j := '8'; j >= i+1; j-- {
			z01.PrintRune(i)
			z01.PrintRune(j)
			if i != '7' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}

	z01.PrintRune('\n')
}
