package piscine

import (
	"sort"
)

func Median(a, b, c, d, e int) int {
	integers := []int{a, b, c, d, e}
	sort.Ints(integers)
	median := integers[2]

	return median
}

func main() {
	a, b, c, d, e := 5, 2, 7, 1, 3
	median := Median(a, b, c, d, e)
	println("Median:", median)
}
