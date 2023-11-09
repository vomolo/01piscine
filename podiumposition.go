package piscine

func PodiumPosition(podium [][]string) [][]string {
	leng := len(podium)
	for i := 0; i < leng; i++ {
		podium[i][0] = podium[leng-i-1][0]
	}
	return podium
}
