package piscine

func StrLen(s string) int {
	count := 0
	for _, r := range s {
		if r > -1 {
			count++
		}
	}
	return count
}
