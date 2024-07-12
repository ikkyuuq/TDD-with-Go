package main

// Sum takes a slice of numbers which has any size.
// @params numbers [5]int => Array take a specific capacity
// @params numbers []int => Slice is a dynamic size of capacity starting with empty
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes a multiple slices returning a new slice containing the totals for each slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	// Create a slice that will have a starting capaity as lengthOfNumbers
	// res := make([]int lengthOfNumbers)
	var res []int
	for _, numbers := range numbersToSum {
		res = append(res, Sum(numbers))
	}
	return res
}

// Calculate the totals of the "tails" of each slice except the first one (the head).
func SumAllTails(numbersToSum ...[]int) []int {
	var res []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			res = append(res, 0)
		} else {
			tail := numbers[1:]
			res = append(res, Sum(tail))
		}
	}
	return res
}
