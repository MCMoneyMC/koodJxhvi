package piscine

func AppendRange(min, max int) []int {
	var output []int
	for i := min; i < max; i++ {
		output = append(output, i)
	}
	return output
}
