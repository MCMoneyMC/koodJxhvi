package piscine

func ToLower(s string) string {
	var output string
	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			output += string(r + 32)
		} else {
			output += string(r)
		}
	}
	return output
}
