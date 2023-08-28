package piscine

func NRune(s string, n int) rune {
	if 0 < n && n < len(s)+1 {
		return []rune(s)[n-1]
	}
	return 0
}
