package piscine

func Any(f func(string) bool, a []string) bool {
	for _, t := range a {
		if f(t) {
			return true
		}
	}
	return false
}
