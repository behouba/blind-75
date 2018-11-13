package sorting

import (
	"fmt"
	"math"
)

func Merge(arr []int) (outputArr []int) {
	var leftArr, rigthArr []int
	mid := len(arr) / 2
	leftArr = append(leftArr, arr[:mid]...)
	rigthArr = append(rigthArr, arr[mid:]...)
	leftArr = append(leftArr, math.MaxInt32)
	rigthArr = append(rigthArr, math.MaxInt32)
	fmt.Println("left = ", leftArr, "rigth = ", rigthArr, "len = ", len(arr))

	Li, Ri := 0, 0

	for i := 0; i < len(arr); i++ {
		if leftArr[Li] >= rigthArr[Ri] {
			outputArr = append(outputArr, rigthArr[Ri])
			Ri++
		} else {
			outputArr = append(outputArr, leftArr[Li])
			Li++
		}
	}

	return
}

func MergeSort(arr []int) []int {
	return arr
}
