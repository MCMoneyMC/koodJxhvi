package piscine

import "github.com/01-edu/z01"

func main() {
	for letter := 'z'; letter >= 'a'; letter-- {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
