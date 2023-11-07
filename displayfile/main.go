package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(arg) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	file, _ := os.Open("quest8.txt")
	ay := make([]byte, 15)
	file.Read(ay)
	fmt.Printf(string(ay))
}
