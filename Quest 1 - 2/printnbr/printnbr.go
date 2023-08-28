package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	// Edgecase handling
	if n == -9223372036854775808 {
		n = 922337203685477580
		z01.PrintRune(45)
		PrintNbr(n)
		z01.PrintRune(48 + 8)
	} else {

		if n < 0 {
			z01.PrintRune(45)
			n = -n
		}

		if n < 10 {
			z01.PrintRune(rune(n + 48))
			return
		}

		PrintNbr(n / 10)
		z01.PrintRune(rune(n%10 + 48))
	}
}
