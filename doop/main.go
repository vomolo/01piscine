package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1 := os.Args[1]
	operator := os.Args[2]
	value2 := os.Args[3]

	var result int64

	if operator == "+" {
		result = strToInt(value1) + strToInt(value2)
	} else if operator == "-" {
		result = strToInt(value1) - strToInt(value2)
	} else if operator == "*" {
		result = strToInt(value1) * strToInt(value2)
	} else if operator == "/" {
		if value2 == "0" {
			println("No division by 0")
			return
		}
		result = strToInt(value1) / strToInt(value2)
	} else if operator == "%" {
		if value2 == "0" {
			println("No modulo by 0")
			return
		}
		result = strToInt(value1) % strToInt(value2)
	}

	print(result)
}

func strToInt(s string) int64 {
	var result int64
	neg := false
	for i, c := range s {
		if i == 0 && c == '-' {
			neg = true
		} else if c >= '0' && c <= '9' {
			result = result*10 + int64(c-'0')
		} else {
			return 0
		}
	}
	if neg {
		result *= -1
	}
	return result
}
