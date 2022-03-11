package arraysnslices

func Sum(array [5]int) int {
	var sum int

	for _, number := range array {
		sum += number
	}

	return sum
}
