package main

import (
	"fmt"

	ps "./problems"
	_ "./sorting"
)

func main() {

	arr := []int{2, 5, 6, 2, 34, 12, 34}
	fmt.Println(ps.CheckAddUp(arr, 40))

}
