package main

import "fmt"

func DescendComb() {
	for i := 99; i >= 10; i-- {
		for j := i - 1; j >= 10; j-- {
			fmt.Printf("%02d %02d", i, j)

			if j != 10 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()
}

func main() {
	DescendComb()
}
