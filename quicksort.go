package main

import (
	"fmt"
)

func main() {
	arr := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}

	fmt.Println("arr = ", quickSort(arr, 0, len(arr)-1))
	fmt.Println("sorted arr = ", arr)
}

func partition(arr []int, firstIndex int, lastIndex int) (sarr []int, position int) {
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
	return arr, i + 1
}

func quickSort(arr []int, firstIndex int, lastIndex int) []int {
	if firstIndex < lastIndex {
		_, p := partition(arr, firstIndex, lastIndex)
		quickSort(arr, firstIndex, p-1)
		quickSort(arr, p+1, lastIndex)
		return arr
	}
	return arr
}
