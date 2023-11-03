package piscine

import (
	"strings"
)

func ConcatParams(args []string) string {
	result := strings.Join(args, "\n")
	return result
}
