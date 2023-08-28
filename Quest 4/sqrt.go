package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	return root(1, nb)
}

func root(guess, nb int) int {
	if guess*guess == nb {
		return guess
	}
	if guess*guess < nb {
		return root(guess+1, nb)
	}
	return 0
}
