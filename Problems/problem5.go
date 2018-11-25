package problems

import (
	"strconv"
	"strings"
)

func PairImpair(arr []int) (pairs []int, impairs []int) {

	for _, x := range arr {
		if x%2 == 0 {
			pairs = append(pairs, x)
		} else {
			impairs = append(impairs, x)
		}
	}
	return
}

func Heureux(n int) bool {
	str := strings.Split(strconv.Itoa(n), "")
	sum := 0
	for _, x := range str {
		nb, _ := strconv.Atoi(x)
		sum += nb * nb
	}

	if sum != 1 {
		if sum == 4 {
			return false
		}
		return Heureux(sum)
	}
	return true
}

// cons(a, b) constructs a pair, and car(pair) and
// cdr(pair) returns the first and last element of that pair.
// For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

// Given this implementation of cons:

// def cons(a, b):
//     def pair(f):
//         return f(a, b)
//     return pair
// Implement car and cdr.
