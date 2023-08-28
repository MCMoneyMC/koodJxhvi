package piscine

import "github.com/01-edu/z01"

func main() {
	for number := '0'; number <= '9'; number++ {
		z01.PrintRune(number)
	}
	z01.PrintRune('\n')
}
