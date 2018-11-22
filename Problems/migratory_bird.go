package problems

type birdType struct {
	Number  int
	Counter int
}

// MigratoryBirds challenge from hackerrank
// Complete the migratoryBirds function below.
func MigratoryBirds(arr []int32) int32 {

	birdsTypes := []birdType{birdType{1, 0}, birdType{2, 0}, birdType{3, 0}, birdType{4, 0}, birdType{5, 0}}

	for _, x := range arr {
		switch x {
		case 1:
			birdsTypes[0].Counter++
			break
		case 2:
			birdsTypes[1].Counter++
			break
		case 3:
			birdsTypes[2].Counter++
			break
		case 4:
			birdsTypes[3].Counter++
			break
		case 5:
			birdsTypes[4].Counter++
		}
	}

	var frequentType birdType
	max := 0
	for _, x := range birdsTypes {
		if x.Counter > max {
			max = x.Counter
			frequentType = x
		}
	}

	return int32(frequentType.Number)
}
