package piscine

func IsAlpha(s string) bool {
	for _, r := range s {
		if (r < 'a' || 'z' < r) && (r < 'A' || 'Z' < r) && (r < '0' || '9' < r) {
			return false
		}
	}
	return true
}
