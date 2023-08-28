package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	sortParams(arr)
	for _, elem := range arr {
		for _, r := range elem {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func sortParams(s []string) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				// Swap elements
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
