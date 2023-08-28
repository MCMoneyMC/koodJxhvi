package piscine

func Capitalize(s string) string {
	arr := []rune(s)
	for i, r := range arr {
		if i == 0 {
			if isLower(r) {
				arr[i] = rune(r - 32)
			}
		} else {
			previous := arr[i-1]
			if isAlphaNumeric(r) {
				if isAlphaNumeric(previous) {
					if isUpper(r) {
						arr[i] = rune(r + 32)
					}
				} else {
					if isLower(r) {
						arr[i] = rune(r - 32)
					}
				}
			}
		}
	}
	return string(arr)
}

func isLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isNumeric(r rune) bool {
	return '0' <= r && r <= '9'
}

func isAlpha(r rune) bool {
	return isLower(r) || isUpper(r)
}

func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || isNumeric(r)
}
