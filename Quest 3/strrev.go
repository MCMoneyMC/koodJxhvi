package piscine

func StrRev(s string) string {
	var output string
	for _, r := range s {
		output = string(r) + output
	}
	return output
}
