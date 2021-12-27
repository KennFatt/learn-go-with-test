package arraysandslices

func Sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return
}

func SumAll(numbersToSum ...[]int) []int {
	sliceCount := len(numbersToSum)

	// make() will allocate new memory to reflect type of `[]int` with
	// 	the size of `sliceCount`
	// It only works for slice, map, and chan only.
	sums := make([]int, sliceCount)

	for i := 0; i < sliceCount; i++ {
		currentNumbers := numbersToSum[i]
		numbersLen := len(currentNumbers)

		if numbersLen == 0 {
			sums[i] = 0
			continue
		}

		sums[i] = Sum(numbersToSum[i])
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for i := 0; i < len(numbersToSum); i++ {
		currentNumbers := numbersToSum[i]
		numbersLen := len(currentNumbers)

		if numbersLen <= 1 {
			sums = append(sums, 0)
			continue
		}

		tails := currentNumbers[1:]
		sums = append(sums, Sum(tails))
	}

	return sums
}
