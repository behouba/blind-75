package main

import "fmt"

func longestPalindrome(s string) string {

	n := len(s)
	maxLen := 0
	res := s[:1]

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			substr := s[i : j+1]
			subLen := len(substr)
			if s[i] == s[j] && isPalindrome(substr) && maxLen < subLen {
				maxLen = subLen
				res = substr
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-(i+1)] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("babaavivad"))
}
