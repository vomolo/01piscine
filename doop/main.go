package main

import (
	"os"
)

func isOperatorValid(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/" || op == "%"
}

func doOperation(value1, value2 int, operator string) (result int, err string) {
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			err = "No division by 0"
		} else {
			result = value1 / value2
		}
	case "%":
		if value2 == 0 {
			err = "No modulo by 0"
		} else {
			result = value1 % value2
		}
	}
	return result, err
}

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}

	value1 := 0
	value2 := 0
	operator := args[2]

	_, err1 := scan(args[1], &value1)
	_, err2 := scan(args[3], &value2)

	if !isOperatorValid(operator) || err1 != "" || err2 != "" {
		return
	}

	result, err := doOperation(value1, value2, operator)

	if err != "" {
		return
	}

	print(result)
}

func scan(s string, x *int) (ok bool, err string) {
	*x = 0
	if s == "" {
		return
	}
	neg := 1
	if s[0] == '-' {
		neg = -1
		s = s[1:]
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return
		}
		*x = *x*10 + int(ch-'0')
		if *x < 0 {
			*x = 0
			err = "Overflow"
			return
		}
	}
	*x *= neg
	ok = true
	return
}
