package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, arg := range os.Args {
		if i > 0 {
			for _, r := range arg {
				z01.PrintRune(r)
			}
			z01.PrintRune('\n')
		}
	}
}
