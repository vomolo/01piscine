package main

import "fmt"

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversed := make([]string, length)

	for i := 0; i < length; i++ {
		reversed[i] = menu[length-i-1]
	}

	return reversed
}

func main() {
	result := ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"})
	fmt.Println(result)
}
