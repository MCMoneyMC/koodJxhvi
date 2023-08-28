package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	temp := QuickSortFunc(a, f)
	copy(a, temp)
}

func QuickSortFunc(arr []string, f func(a, b string) int) []string {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)-1]
	var smaller []string
	var bigger []string
	for index, s := range arr {
		if index != len(arr)-1 {
			if f(pivot, s) > 0 {
				smaller = append(smaller, s)
			} else {
				bigger = append(bigger, s)
			}
		}
	}
	output := append(QuickSortFunc(smaller, f), pivot)
	output = append(output, QuickSortFunc(bigger, f)...)
	return output
}
