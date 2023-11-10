package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	i := 0
	for i < len(str) {
		word := ""
		for j := 0; j < 5; j++ {
			if i+j < len(str) {
				if str[i+j] != ' ' {
					word += string(str[i+j])
				} else {
					j--
				}
			} else {
				break
			}
		}
		result += word
		i += len(word) + 1
	}

	return result + "\n"
}
