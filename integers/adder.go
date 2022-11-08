package integers

// Returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Returns the sum of zero or more integers.
func AddAll(nums ...int) int {
	var result int
	for _, n := range nums {
		result += n
	}

	return result
}
