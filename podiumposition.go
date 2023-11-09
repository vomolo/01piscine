package piscine

import "sort"

func PodiumPosition(podium [][]string) [][]string {
	// Sort the podium slice in ascending order
	sort.Slice(podium, func(i, j int) bool {
		return podium[i][0] < podium[j][0]
	})

	return podium
}
