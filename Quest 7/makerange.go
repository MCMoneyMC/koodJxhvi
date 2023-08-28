package piscine

func MakeRange(min, max int) []int {
	if max-min < 1 {
		return nil
	}
	output := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		output[i] = min + i
	}
	return output
}
