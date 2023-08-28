package piscine

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)-1]
	var smaller []int
	var bigger []int
	for index, i := range arr {
		if index != len(arr)-1 {
			if i <= pivot {
				smaller = append(smaller, i)
			} else {
				bigger = append(bigger, i)
			}
		}
	}
	output := append(QuickSort(smaller), pivot)
	output = append(output, QuickSort(bigger)...)
	return output
}
