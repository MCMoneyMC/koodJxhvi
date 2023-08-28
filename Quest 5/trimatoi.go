package piscine

func TrimAtoi(s string) int {
	if !isValidString(s) {
		return 0
	}
	output := 0
	isNegative := false
	for _, r := range s {
		if r == '-' && output == 0 {
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
	for _, r := range s {
		if '0' <= r && r <= '9' {
			count++
		}
	}
	return count > 0
}
