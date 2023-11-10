package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	operator := os.Args[2]
	value2, err2 := strconv.ParseInt(os.Args[3], 10, 64)

	if err1 != nil || err2 != nil {
		return
	}

	result := int64(0)
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = value1 % value2
	default:
		return
	}

	fmt.Println(result)
}
