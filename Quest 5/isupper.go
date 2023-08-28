package piscine

func IsUpper(s string) bool {
	for _, r := range s {
		if r < 'A' || 'Z' < r {
			return false
		}
	}
	return true
}
