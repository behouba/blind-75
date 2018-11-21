package problems

// Disible Sum Pairs challenge from hackerranck
// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32
	for i, x := range ar {
		for _, y := range ar[i+1:] {
			if ((x + y) % k) == 0 {
				count++
			}
		}
	}
	return count
}
