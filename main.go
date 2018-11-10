package main

import (
	"fmt"

	"./sorting"
)

func main() {
	arr := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}

	fmt.Println("sorted arr = ", sorting.QuickSort(arr, 0, len(arr)-1))
}
