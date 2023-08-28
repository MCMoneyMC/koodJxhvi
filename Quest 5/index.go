package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	orig := []rune(s)
	pattern := []rune(toFind)

	for i := 0; i < len(orig); i++ {
		for j := 0; j < len(pattern); j++ {
			if orig[i+j] != pattern[j] {
				break
			}
			if j == len(pattern)-1 {
				return i
			}
		}
	}

	return -1
}
