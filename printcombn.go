package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	digits := make([]rune, n)
	printCombinations(digits, 0, '0', n)
}

func printCombinations(digits []rune, index int, startChar rune, n int) {
	if index == n {
		z01.PrintRune(digits[0])
		for i := 1; i < n; i++ {
			z01.PrintRune(',')
			z01.PrintRune(' ')
			z01.PrintRune(digits[i])
		}
		z01.PrintRune('\n')
		return
	}

	for i := startChar; i <= '9'; i++ {
		digits[index] = i
		printCombinations(digits, index+1, i+1, n)
	}
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
}
