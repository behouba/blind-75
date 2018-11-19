package problems

import (
	"fmt"
)

/*Good morning! Here's your coding interview problem for today.

This problem was recently asked by Google.

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?*/

// CheckAddUp is my first solution to this problem
func CheckAddUp(arr []int, numb int) bool {
	for i := 0; i < len(arr); i++ {
		k := i + 1
		for j := k; j < len(arr); j++ {
			fmt.Println(arr[i], " ", arr[j])
			if arr[i]+arr[j] == numb {
				fmt.Println("numb are =", arr[i], " and ", arr[j])
				return true
			}
		}
	}
	return false
}
