package piscine

func MakeRange(min int, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	res := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = min + i
	}
	return res
}
