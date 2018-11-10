package sorting

// partitioning function which return the pivot index after partion
func partition(arr []int, firstIndex int, lastIndex int) (position int) {
	pivot := arr[lastIndex]
	i := firstIndex - 1
	j := firstIndex
	for j <= lastIndex-1 {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
		j++
	}
	arr[lastIndex], arr[i+1] = arr[i+1], arr[lastIndex]
	return i + 1
}

// quicksort with call to partition function and recursives calls
func QuickSort(arr []int, firstIndex int, lastIndex int) []int {
	if firstIndex < lastIndex {
		p := partition(arr, firstIndex, lastIndex)
		QuickSort(arr, firstIndex, p-1)
		QuickSort(arr, p+1, lastIndex)
		return arr
	}
	return arr
}
