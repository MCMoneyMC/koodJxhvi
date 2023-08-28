package piscine

func ToUpper(s string) string {
	var output string
	for _, r := range s {
		if 'a' <= r && r <= 'z' {
			output += string(r - 32)
		} else {
			output += string(r)
		}
	}
	return output
}
