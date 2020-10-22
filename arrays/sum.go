package arrays

// Sum calculates the total from a slice of numbers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll calculates the total of each slice
func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails calculates the total of each slice except the first number
func SumAllTails(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		if len(numbers) < 2 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
