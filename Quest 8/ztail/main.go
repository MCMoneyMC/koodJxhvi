package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	giveError := false
	if len(args) > 3 {
		// Read each file and print its contents to stdout
		for i, arg := range args[3:] {
			text, err := os.ReadFile(arg)
			if err != nil {
				fmt.Printf("%s\n", err)
				giveError = true
			} else {
				if len(args) > 4 {
					if i > 0 {
						fmt.Printf("\n")
					}
					fmt.Printf("==> %s <==\n", arg)
				}
				from := len(text) - Atoi(args[2])
				if from > -1 {
					fmt.Printf("%s", text[len(text)-Atoi(args[2]):])
				} else {
					fmt.Printf("%s", text)
				}
			}
		}
	}
	if giveError {
		os.Exit(1)
	}
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
