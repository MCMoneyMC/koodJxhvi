package piscine

func Atoi(s string) int {
	if !isValidString(s) {
		return 0
	}
	output := 0
	isNegative := false
	for _, r := range s {
		if r == '-' {
			isNegative = true
		}
		if isNegative {
			if '0' <= r && r <= '9' {
				output *= 10
				output -= (int(r) - 48)
			}
		} else {
			if '0' <= r && r <= '9' {
				output *= 10
				output += (int(r) - 48)
			}
		}
	}
	return output
}

func isValidString(s string) bool {
	count := 0
	for i, r := range s {
		if r == '-' || r == '+' {
			if i != 0 {
				return false
			}
		} else {
			if !('0' <= r && r <= '9') {
				return false
			}
		}
		if count > 1 {
			return false
		}
	}
	return true
}
