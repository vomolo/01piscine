package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i != '7' {
					zo1.PrintRune(',')
					zo1.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
