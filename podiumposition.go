package piscine

func PodiumPosition(podium [][]string) [][]string {
	// Create an array of slices to represent the positions
	positions := [4][]string{}

	// Iterate over the input podium and place competitors in their respective positions
	for _, competitor := range podium {
		position := getPosition(competitor)
		positions[position-1] = append(positions[position-1], competitor[0])
	}

	// Convert the array of slices to a slice of slices
	result := make([][]string, 0, 4)
	for i := 0; i < 4; i++ {
		if len(positions[i]) > 0 {
			result = append(result, positions[i])
		}
	}

	return result
}

func getPosition(competitor []string) int {
	if len(competitor) != 1 {
		return 0
	}

	switch competitor[0] {
	case "1st Place":
		return 1
	case "2nd Place":
		return 2
	case "3rd Place":
		return 3
	case "4th Place":
		return 4
	default:
		return 0
	}
}
