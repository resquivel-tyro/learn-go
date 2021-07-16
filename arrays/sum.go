package arrays

func SumArray(numbers [5]int) int {
	sum := 0
	// _ in this case is the index value, which we chose to ignore
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	// _ in this case is the index value, which we chose to ignore
	for _, number := range numbers {
		sum += number
	}
	return sum
}
