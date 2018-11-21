package problems

import (
	"fmt"
)

// bonAppeti challenge from hackerrank solved 21/11/2018
// Complete the bonAppetit function below.
func bonAppetit(bill []int32, k int32, b int32) {
	var regSum int32
	var regPar int32
	for _, x := range bill {
		regSum += x
	}

	regSum = regSum - bill[k]
	regPar = regSum / 2

	if regPar == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - regPar)
	}
}
