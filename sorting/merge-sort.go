package sorting

import (
	"math"
)

// Merge function merge two sorted part of an array and merge them
func Merge(arr []int) []int {
	var leftArr, rigthArr []int
	mid := len(arr) / 2
	leftArr = append(leftArr, arr[:mid]...)
	rigthArr = append(rigthArr, arr[mid:]...)
	leftArr = append(leftArr, math.MaxInt32)
	rigthArr = append(rigthArr, math.MaxInt32)

	Li, Ri := 0, 0

	for i := 0; i < len(arr); i++ {
		if leftArr[Li] >= rigthArr[Ri] {
			arr[i] = rigthArr[Ri]
			Ri++
		} else {
			arr[i] = leftArr[Li]
			Li++
		}
	}
	return arr
}

// MergeSort function through recursive
// call and merge with Merge function to sort an array

func MergeSort(arr []int) []int {

	// save the length off the array
	arrLength := len(arr)

	// call MergeSort until we get only two element in the array
	if arrLength > 1 {
		mid := arrLength / 2
		MergeSort(arr[:mid])
		MergeSort(arr[mid:])
		Merge(arr)
	}
	return arr
}
