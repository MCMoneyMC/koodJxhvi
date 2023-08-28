package piscine

func IsLower(s string) bool {
	for _, r := range s {
		if r < 'a' || 'z' < r {
			return false
		}
	}
	return true
}
