package piscine

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || '9' < r {
			return false
		}
	}
	return true
}
