package piscine

func Unmatch(arr []int) int {
	for i, char := range arr {
		count := 1
		for j, ch := range arr {
			if char == ch && i != j {
				count++
			}
		}
		if count%2 == 1 {
			return char
		}
	}
	return -1
}
