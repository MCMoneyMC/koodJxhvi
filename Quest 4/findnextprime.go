package piscine

func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}
	if nb%2 == 0 {
		if PrimeCalculation(nb+1, 2) {
			return nb + 1
		}
		return FindNextPrime(nb + 3)
	}
	if PrimeCalculation(nb, 2) {
		return nb
	}
	return FindNextPrime(nb + 2)
}

func PrimeCalculation(nb, divisor int) bool {
	if nb < 2 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%divisor == 0 {
		return false
	}
	if divisor*divisor > nb {
		return true
	}
	return PrimeCalculation(nb, divisor+1)
}
