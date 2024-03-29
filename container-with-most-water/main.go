package main

// Problem: https://leetcode.com/problems/container-with-most-water/

// Naive solution, O(n^2) time complexity
func bruteForceSolution(height []int) int {
	vol := 0
	n := len(height)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v := min(height[i], height[j]) * (j - i)
			if v > vol {
				vol = v
			}
		}
	}

	return vol
}

// Solution based on hint 2: O(n)
func solution(height []int) int {
	i := 0
	j := len(height) - 1
	vol := 0

	for i < j {

		v := min(height[i], height[j]) * (j - i)
		if v > vol {
			vol = v
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}

	return vol
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//
}
