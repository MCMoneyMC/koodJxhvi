package piscine

func Split(s, sep string) []string {
	var output []string
	length := len(sep)
	seps := Index(s, sep)
	if seps[0] != 0 {
		s = sep + s
		for i := 0; i < len(seps); i++ {
			seps[i] = seps[i] + length
		}
		seps = append([]int{0}, seps...)
	}
	seps = append(seps, len(s))
	for i := 0; i < len(seps)-1; i++ {
		if seps[i]+length != seps[i+1] {
			output = append(output, s[seps[i]+length:seps[i+1]])
		}
	}
	return output
}

func Index(s string, toFind string) []int {
	if toFind == "" {
		return nil
	}
	orig := []rune(s)
	pattern := []rune(toFind)
	var matches []int

	for i := 0; i < len(orig); i++ {
		for j := 0; j < len(pattern); j++ {
			if orig[i+j] != pattern[j] {
				break
			}
			if j == len(pattern)-1 {
				matches = append(matches, i)
			}
		}
	}

	return matches
}
