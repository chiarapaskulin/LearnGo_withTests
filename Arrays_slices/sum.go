package arrays

func Sum(numbers []int) (sum int) {

	/* V1:
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	*/

	for _, numbers := range numbers {
		sum += numbers
	}

	return sum

}

func SumAll(numbersToSum ...[]int) []int {
	/* V1:
	sums := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	*/

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
