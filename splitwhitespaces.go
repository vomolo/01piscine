package piscine

func SplitWords(input string) []string {
	var words []string
	word := ""
	for _, char := range input {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}
