package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help()
		return
	}
	getFlags(args)
}

func getFlags(args []string) {
	order := false
	used := false
	var s string
	var insert string
	for _, arg := range args {
		if len(arg) > 9 {
			if arg[:9] == "--insert=" {
				insert = arg[9:]
				continue
			}
		}
		if len(arg) >= 7 {
			if arg[:7] == "--order" {
				order = true
				continue
			}
		}
		if len(arg) > 3 {
			if arg[:3] == "-i=" {
				insert = arg[3:]
				continue
			}
		}
		if len(arg) >= 2 {
			if arg[:2] == "-o" {
				order = true
				continue
			}
		}
		if len(arg) > 0 {
			if arg[:1] == "-" {
				help()
				return
			}
		}
		if !used {
			s = arg
		} else {
			help()
			return
		}
	}
	s += insert
	if order {
		s = sort([]rune(s))
	}
	printOutput(s)
}

func help() {
	s := "--insert\n"
	s += "  -i\n"
	s += "	 This flag inserts the string into the string passed as argument.\n"
	s += "--order\n"
	s += "  -o\n"
	s += "	 This flag will behave like a boolean, if it is called it will order the argument."
	printOutput(s)
}

func printOutput(s string) {
	fmt.Println(s)
}

func sort(s []rune) string {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				// Swap elements
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return string(s)
}
