package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	arr := convertToSlice(n)
	sortSlice(arr)
	for _, r := range arr {
		z01.PrintRune(rune(int(r) + 48))
	}
}

func convertToSlice(n int) []int {
	var output []int
	for {
		if n < 10 {
			output = append(output, n)
			break
		}
		output = append(output, n%10)
		n /= 10
	}
	return output
}

func sortSlice(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
