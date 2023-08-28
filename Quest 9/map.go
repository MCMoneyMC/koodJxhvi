package piscine

func Map(f func(int) bool, a []int) []bool {
	var output []bool
	for _, nbr := range a {
		output = append(output, f(nbr))
	}
	return output
}
