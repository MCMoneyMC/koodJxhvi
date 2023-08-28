package piscine

import "github.com/01-edu/z01"

func main() {
	string := "abcdefghijklmnopqrstuvwxyz"
	for _, runes := range []rune(string) {
		z01.PrintRune(runes)
	}
	z01.PrintRune('\n')
}
