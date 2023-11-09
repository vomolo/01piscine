package piscine

func PodiumPosition(podium [][]string) [][]string {
	positionMap := make(map[string]int)

	for i, place := range podium {
		if len(place) == 1 {
			positionMap[place[0]] = i
		}
	}

	competitors := make([]string, len(podium))

	for place, pos := range positionMap {
		competitors[pos] = place
	}

	result := make([][]string, len(podium))
	for i, place := range competitors {
		result[i] = []string{place}
	}

	return result
}
