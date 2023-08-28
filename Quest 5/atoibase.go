package piscine

func AtoiBase(s string, base string) int {
	if !(isBaseValid([]rune(base))) {
		return 0
	}
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

func isBaseValid(base []rune) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
		if base[i] == '+' || base[i] == '-' {
			return false
		}
	}
	return true
}
