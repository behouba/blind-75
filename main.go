package main

import (
	"fmt"

	ps "./problems"
	_ "./sorting"
)

func main() {

	arr := []int32{1, 4, 4, 4, 5, 3}
	fmt.Println(ps.MigratoryBirds(arr))

}
