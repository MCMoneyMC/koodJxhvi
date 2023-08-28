package piscine

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	findPositions(nil)
}

func findPositions(currentPosition []rune) {
	if len(currentPosition) == 8 {
		printRow(currentPosition)
	} else {
		for i := '1'; i < '9'; i++ {
			if checkValid(i, currentPosition) {
				currentPosition := append(currentPosition, i)
				findPositions(currentPosition)
			}
		}
	}
}

func checkValid(element rune, position []rune) bool {
	column := 0
	for _, queen := range position {
		column += 1
		if queen == element {
			return false
		}
		if int(element-queen) == len(position)+1-column || int(queen-element) == len(position)+1-column {
			return false
		}
	}
	return true
}

func printRow(row []rune) {
	for _, r := range row {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// first execution 2.958426ms
