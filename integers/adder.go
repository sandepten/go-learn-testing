package integers

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

// Subtract takes two integers and returns the difference of them
func Subtract(x, y int) int {
	return x - y
}

// ArraySum takes an array of integers and returns the sum of them
func ArraySum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// ArraySumAll takes a varying number of slices, returning a new slice containing the totals for each slice passed in
func ArraySumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, ArraySum(numbers))
	}
	return sums
}
