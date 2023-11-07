package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}
	var asc, desc bool
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			desc = true
		} else if f(a[i-1], a[i]) < 0 {
			asc = true
		}
		if asc && desc {
			return false
		}
	}
	return true
}
