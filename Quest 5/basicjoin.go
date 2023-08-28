package piscine

func BasicJoin(elems []string) string {
	var output string
	for _, s := range elems {
		output += s
	}
	return output
}
