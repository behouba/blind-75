package problems

import (
	"fmt"
)

// Birthday chocolate challenge from hackerrank
func Birthday(s []int32, d int32, m int32) int32 {

	if len(s) < 0 {
		return 0
	}

	for i, x := range s {
		var chocoBar []int32
		if x == d && m == 1 {
			return Birthday(s[i+1:], d, m) + 1
		}
		chocoBar = append(chocoBar, x)
		for j, y := range s[i+1:] {

			chocoBar = append(chocoBar, y)
			if arrSum(chocoBar) == d && len(chocoBar) == int(m) {
				fmt.Println(j, i)
				return Birthday(s[i+1:], d, m) + 1
			}

			if arrSum(chocoBar) > d || len(chocoBar) > int(m) {
				break
			}

		}

	}
	return 0
}

func arrSum(arr []int32) int32 {
	var output int32
	for _, x := range arr {
		output += x
	}
	return output
}
