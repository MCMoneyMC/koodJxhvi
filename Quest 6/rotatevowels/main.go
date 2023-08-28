package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printOutput("")
		return
	}
	var s string
	for i, arg := range args {
		s += arg
		if i < len(args)-1 && len(arg) != 0 {
			s += " "
		}
	}
	vowelLocations := getLocations(s)
	swapped := swapVowels(s, vowelLocations)
	printOutput(swapped)
}

func getLocations(s string) []int {
	var output []int
	for i, r := range s {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
			output = append(output, i)
		}
	}
	return output
}

func swapVowels(s string, indexes []int) string {
	output := []rune(s)
	for i := 0; i < len(indexes)/2; i++ {
		output[indexes[i]], output[indexes[len(indexes)-1-i]] = output[indexes[len(indexes)-1-i]], output[indexes[i]]
	}
	return string(output)
}

func printOutput(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
