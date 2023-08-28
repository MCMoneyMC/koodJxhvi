package piscine

func BasicAtoi2(s string) int {
	if !isValidString(s) {
		return 0
	}
	output := 0
	for _, r := range s {
		if '0' <= r && r <= '9' {
			output *= 10
			output += (int(r) - 48)
		}
	}
	return output
}

func isValidString(s string) bool {
	for _, r := range s {
		if !(('0' <= r && r <= '9') || r == '-') {
			return false
		}
	}
	return true
}
