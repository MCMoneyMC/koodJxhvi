package piscine

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') {
			count++
		}
	}
	return count
}
