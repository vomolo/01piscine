package main

import (
	"os"
)

func PrintConsole(str string) {
	os.Stdout.WriteString(str)
	os.Stdout.WriteString("\n")
}

// ... (other functions)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		PrintConsole("0")
		return
	}

	if !(IsNumeric(args[0]) && IsNumeric(args[2])) {
		PrintConsole("0")
		return
	}

	funcsMap := map[string]func(string, string){
		"+": Plus,
		"-": Deduct,
		"/": Devide,
		"*": Multiply,
		"%": Mod,
	}

	operator := args[1]
	if function, exists := funcsMap[operator]; exists {
		function(args[0], args[2])
	} else {
		PrintConsole("0")
	}
}
