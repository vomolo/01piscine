package piscine

func ForEach(f func(int), a []int) {
	for _, char := range a {
		f(char)
	}
}
