package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	output := 1
	for i := 0; i < nb; i++ {
		output *= i + 1
	}
	return output
}
