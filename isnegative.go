package main

import "fmt"

func isNegative(number int) {
    if number < 0 {
        fmt.Println("T")
    } else {
        fmt.Println("F")
    }
}

func main() {
    isNegative(-7) // This will print 'T'
    isNegative(11)  // This will print 'F'
}
