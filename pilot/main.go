package piscine

import "fmt"

const AIRCRAFT1 = 1

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}
