package problems

// breakingRecords challenge from hackerranck
// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {
	res := []int32{0, 0}
	max := scores[0]
	min := scores[0]

	for i := 0; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
			res[0] = res[0] + 1
		} else if scores[i] < min {
			min = scores[i]
			res[1]++
		}
	}
	return res
}
