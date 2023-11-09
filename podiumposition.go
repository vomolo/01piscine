package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium)
	for i := 0; i < length; i++ {
		podium[i][0] = podium[length-i-1][0]
	}
	return podium
}
