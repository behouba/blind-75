package main

// Problem: https://leetcode.com/problems/container-with-most-water/

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
