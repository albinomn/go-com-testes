package arraysnslices

func Sum(slice []int) int {
	var sum int

	for _, number := range slice {
		sum += number
	}

	return sum
}

func SumAll(slices ...[]int) (sum []int) {
	for _, numbers := range slices {
		sum = append(sum, Sum(numbers))
	}

	return
}
