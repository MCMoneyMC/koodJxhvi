package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	if args[0] == "--upper" {
		translate(args[1:], 64)
	} else {
		translate(args, 96)
	}
}

func translate(elems []string, beforeFirst int) {
	for _, i := range elems {
		temp := Atoi(i)
		if 0 < temp && temp < 27 {
			z01.PrintRune(rune(temp + beforeFirst))
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	if !isValidString(s) {
		return 0
	}
	output := 0
	isNegative := false
	for _, r := range s {
		if r == '-' {
			isNegative = true
		}
		if isNegative {
			if '0' <= r && r <= '9' {
				output *= 10
				output -= (int(r) - 48)
			}
		} else {
			if '0' <= r && r <= '9' {
				output *= 10
				output += (int(r) - 48)
			}
		}
	}
	return output
}

func isValidString(s string) bool {
	count := 0
	for i, r := range s {
		if r == '-' || r == '+' {
			if i != 0 {
				return false
			}
		} else {
			if !('0' <= r && r <= '9') {
				return false
			}
		}
		if count > 1 {
			return false
		}
	}
	return true
}
