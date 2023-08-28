package piscine

func Join(elems []string, sep string) string {
	var output string
	for i, s := range elems {
		output += s
		if i < len(elems)-1 {
			output += sep
		}
	}
	return output
}
