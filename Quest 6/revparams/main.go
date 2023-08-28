package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	reverse(arguments)
	for i, arg := range arguments {
		if i < len(arguments)-1 {
			for _, r := range arg {
				z01.PrintRune(r)
			}
			z01.PrintRune('\n')
		}
	}
}

func reverse(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
