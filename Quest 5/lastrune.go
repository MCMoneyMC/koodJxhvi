package piscine

func LastRune(s string) rune {
	var output rune
	for _, r := range s {
		output = r
	}
	return output
}
