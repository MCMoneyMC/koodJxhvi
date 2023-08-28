package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	for _, r := range "x = " {
		z01.PrintRune(r)
	}
	PrintNbr(points.x)
	for _, r := range ", y = " {
		z01.PrintRune(r)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}

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
