package piscine

func ConcatParams(args []string) string {
	var output string
	for i, s := range args {
		output += s
		if i < len(args)-1 {
			output += "\n"
		}
	}
	return output
}
