package sorting

func InsertionSort(arr []int) []int {

	arrLength := len(arr)

	for j := 1; j < arrLength; j++ {

		key := arr[j]
		i := j - 1

		for i > -1 && key < arr[i] {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}
	return arr
}
