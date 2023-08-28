package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}
	if !(isNumeric(args[1]) && isNumeric(args[3]) && isOperator(args[2])) {
		return
	}
	a := Atoi(args[1])
	b := Atoi(args[3])
	var output int
	switch getOperator(args[2]) {
	case 1:
		output = a + b
		if a < 0 && b < 0 {
			if output > a || output > b {
				return
			}
		} else if a < 0 || b < 0 {
			if output > a && output > b {
				return
			}
		} else {
			if output < a || output < b {
				return
			}
		}
	case 2:
		output = a - b
		if output > a {
			return
		}
	case 3:
		output = a * b
		if a < 0 && b < 0 {
			if output < a || output < b {
				return
			}
		} else if a < 0 || b < 0 {
			if output > a && output > b {
				return
			}
		} else {
			if output < a || output < b {
				return
			}
		}
	case 4:
		if b == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		output = a / b
	case 5:
		if b == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		output = a % b
	}
	s := PrintNbr(output, "")
	os.Stdout.Write([]byte(s + "\n"))
}

func isNumeric(s string) bool {
	used := false
	for _, r := range s {
		if !('0' <= r && r <= '9') {
			if r == '-' && !used {
				used = true
			} else {
				return false
			}
		}
	}
	return true
}

func isOperator(s string) bool {
	if (len(s)) != 1 {
		return false
	}
	if rune(s[0]) == '+' || rune(s[0]) == '/' || rune(s[0]) == '*' || rune(s[0]) == '-' || rune(s[0]) == '%' {
		return true
	}
	return false
}

func getOperator(s string) int {
	switch rune(s[0]) {
	case '+':
		return 1
	case '-':
		return 2
	case '*':
		return 3
	case '/':
		return 4
	}
	return 5
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

func PrintNbr(n int, s string) string {
	// Edgecase handling
	if n == -9223372036854775808 {
		return "-9223372036854775808"
	}

	if n < 0 {
		return "-" + PrintNbr(-n, s)
	}

	if n < 10 {
		s += string(rune(n + 48))
		return s
	}

	s = PrintNbr(n/10, s)
	s += string(rune(n%10 + 48))
	return s
}
