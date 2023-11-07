package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func PrintString(st string) {
	arraystr := []rune(st)

	for i := range arraystr {
		z01.PrintRune(arraystr[i])
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintInteger(numb int) {
	numXY := '0'
	if numb/10 > 0 {
		PrintInteger(numb / 10)
	}
	for i := 0; i < numb%10; i++ {
		numXY++
	}
	z01.PrintRune(numXY)
}

func main() {
	var points point

	setPoint(&points)

	PrintString("x = ")
	PrintInteger(points.x)
	PrintString(", y = ")
	PrintInteger(points.y)
	z01.PrintRune('\n')
}
