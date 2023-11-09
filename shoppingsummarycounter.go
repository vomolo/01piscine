package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	Summary := make(map[string]int)
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != '' {
		word += string(str[i])
		} else {
		Summary[word]++
		word = ""
		}
	}
	Summary[word]++
	return Summary
}
