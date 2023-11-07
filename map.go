package piscine

func Map(f func(int) bool, a []int) []bool {
	var ay []bool
	for _, elem := range a {
		ay = append(ay, f(elem))
	}
	return ay
}
