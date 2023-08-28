package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return ConvertToBase(ConvertToNum(nbr, baseFrom), baseTo)
}

func ConvertToNum(s string, base string) int {
	max := len(base)
	output := 0
	for i, r := range s {
		output += index(base, r)
		if i < len(s)-1 {
			output *= max
		}
	}
	return output
}

func index(s string, pattern rune) int {
	for i, r := range s {
		if r == pattern {
			return i
		}
	}
	return -1
}

func ConvertToBase(nbr int, base string) string {
	chars := []rune(base)
	max := len(base)
	var output []rune
	for {
		if nbr < max {
			output = append([]rune{chars[nbr]}, output...)
			break
		} else {
			output = append([]rune{chars[nbr%max]}, output...)
			nbr /= max
		}
	}
	return string(output)
}
