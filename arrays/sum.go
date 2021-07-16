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

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbersSlice := len(numbersToSum)
	sums := make([]int, lengthOfNumbersSlice) // create a slice with starting capacity lengthOfNumbers

	for i, numbersToSum := range numbersToSum {
		sums[i] = SumSlice(numbersToSum)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbersToSum := range numbersToSum {
		tail := numbersToSum[1:]
		sums = append(sums, SumSlice(tail))
	}
	return sums
}