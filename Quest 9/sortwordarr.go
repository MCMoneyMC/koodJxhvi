package piscine

func SortWordArr(a []string) {
	temp := QuickSort(a)
	copy(a, temp)
}

func QuickSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)-1]
	var smaller []string
	var bigger []string
	for index, s := range arr {
		if index != len(arr)-1 {
			if s <= pivot {
				smaller = append(smaller, s)
			} else {
				bigger = append(bigger, s)
			}
		}
	}
	output := append(QuickSort(smaller), pivot)
	output = append(output, QuickSort(bigger)...)
	return output
}
