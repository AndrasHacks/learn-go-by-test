package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	//lengthOfNumbers := len(numbersToSum)
	// Create a slice, with the starting capacity of
	// sums = make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		var tails []int
		for idx, number := range numbers {
			if idx > 0 {
				tails = append(tails, number)
			}
		}
		sums = append(sums, Sum(tails))
	}
	return sums
}
