package main

import (
	"errors"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, op, b := os.Args[1], os.Args[2], os.Args[3]

	result, err := doOperation(a, op, b)
	if err != nil {
		return
	}

	writeString(result)
}

func doOperation(a, op, b string) (string, error) {
	switch op {
	case "+":
		return add(a, b), nil
	case "-":
		return subtract(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		result, err := divide(a, b)
		if err != nil {
			return "", err
		}
		return result, nil
	case "%":
		result, err := modulo(a, b)
		if err != nil {
			return "", err
		}
		return result, nil
	default:
		return "", errors.New("invalid operator")
	}
}

func add(a, b string) string {
	return toString(toInt(a) + toInt(b))
}

func subtract(a, b string) string {
	return toString(toInt(a) - toInt(b))
}

func multiply(a, b string) string {
	return toString(toInt(a) * toInt(b))
}

func divide(a, b string) (string, error) {
	if b == "0" {
		writeString("No division by 0")
		return "", errors.New("division by zero")
	}
	return toString(toInt(a) / toInt(b)), nil
}

func modulo(a, b string) (string, error) {
	if b == "0" {
		writeString("No modulo by 0")
		return "", errors.New("modulo by zero")
	}
	return toString(toInt(a) % toInt(b)), nil
}

func toInt(s string) int {
	result := 0
	neg := false
	for i, c := range s {
		if i == 0 && c == '-' {
			neg = true
			continue
		}
		result = result*10 + int(c-'0')
	}
	if neg {
		return -result
	}
	return result
}

func toString(n int) string {
	if n < 0 {
		return "-" + toString(-n)
	}
	if n < 10 {
		return string('0' + n)
	}
	return toString(n/10) + string('0'+n%10)
}

func writeString(s string) {
	// Use os.Stdout.Write() to write the string to standard output
	os.Stdout.Write([]byte(s))
}
