package piscine

func SplitWhiteSpaces(s string) []string {
	spaces := []int{-1}
	for i, r := range s {
		if r == ' ' || r == '\n' || r == '\t' {
			spaces = append(spaces, i)
		} else if i == len(s)-1 {
			spaces = append(spaces, i+1)
		}
	}
	var output []string
	for i := 0; i < len(spaces)-1; i++ {
		if spaces[i]+1 != spaces[i+1] {
			output = append(output, s[spaces[i]+1:spaces[i+1]])
		}
	}
	return output
}
