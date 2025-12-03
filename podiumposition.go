package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			a := podium[j][0]
			b := podium[j+1][0]
			if isGreater3(a, b) {
				podium[j], podium[j+1] = []string{b}, []string{a}
			}
		}
	}

	return podium
}

func isGreater3(s1, s2 string) bool {
	n := len(s1)
	s2len := len(s2)
	greater := false
	for i := 0; i < n; i++ {
		if i < s2len && s1[i] > s2[i] {
			return true
		}
	}

	if !greater && len(s1) > len(s2) {
		return true
	}

	return greater
}
