package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	chars := []rune(base)
	if !isValidBase([]rune(base)) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	max := len(base)
	var output []rune
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			output = []rune{chars[-(nbr % max)]}
			nbr /= max

		}
		nbr *= -1
	}
	for {
		if nbr < max {
			output = append([]rune{chars[nbr]}, output...)
			break
		} else {
			output = append([]rune{chars[nbr%max]}, output...)
			nbr /= max
		}
	}
	printRow(output)
}

func isValidBase(base []rune) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
		if base[i] == '+' || base[i] == '-' {
			return false
		}
	}
	return true
}

func printRow(s []rune) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
