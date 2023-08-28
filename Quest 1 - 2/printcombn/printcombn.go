package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	PrintCombn_Step(n, '0', "", n)
	z01.PrintRune('\n')
}

func PrintCombn_Step(depth int, starting rune, prefix string, max int) {
	if depth > 1 {
		for i := starting; i <= '9'; i++ {
			PrintCombn_Step(depth-1, i+1, prefix+string(i), max)
		}
	} else {
		for i := starting; i <= '9'; i++ {
			if i != rune(max)+'0'-1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			for _, r := range prefix {
				z01.PrintRune(r)
			}
			z01.PrintRune(i)
		}
		if max != 1 {
			prefix = prefix[:len(prefix)-1]
		}
	}
}
