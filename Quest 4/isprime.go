package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	return PrimeCalculation(nb, 2)
}

func PrimeCalculation(nb, divisor int) bool {
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
