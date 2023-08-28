package piscine

func FirstRune(s string) rune {
	var output rune
	for _, r := range s {
		output = r
		break
	}
	return output
}
