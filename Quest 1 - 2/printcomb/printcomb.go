package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var runes []rune
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				runes = append(runes, i, j, k)
				if i != '7' {
					runes = append(runes, ',', ' ')
				} else {
					runes = append(runes, '\n')
				}
			}
		}
	}

	for _, letter := range runes {
		z01.PrintRune(letter)
	}
}
