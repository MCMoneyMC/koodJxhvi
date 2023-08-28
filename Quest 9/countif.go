package piscine

func CountIf(f func(string) bool, tab []string) int {
	output := 0
	for _, s := range tab {
		if f(s) {
			output++
		}
	}
	return output
}
