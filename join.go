package piscine

func Join(strs []string, sep string) string {
	concat := ""
	for i, jo := range strs {
		concat += jo
		if i != len(strs)-1 {
			concat += sep
		}
	}
	return concat
}
