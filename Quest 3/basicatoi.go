package piscine

func BasicAtoi(s string) int {
	output := 0
	for _, r := range s {
		output *= 10
		output += (int(r) - 48)
	}
	return output
}
