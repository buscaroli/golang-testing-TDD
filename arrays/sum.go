package arrays

import "math/rand"

// Returns the sum of the elements of an array of 10000 ints. For loop using range.
func SumIntArray(nums [10000]int) int {
	var result int

	for _, n := range nums {
		result += n
	}

	return result
}

// Returns the sum of the elements of an array of 5 ints. For loop NOT using range.
func SumIntArray2(nums [10000]int) int {
	var result int

	for i := 0; i < 10000; i++ {
		result += nums[i]
	}

	return result
}

// Returns the sum of the elements of a slice of ints.
func SumIntSlice(nums []int) int {
	var result int

	for _, n := range nums {
		result += n
	}

	return result
}

// Returns an array of 10000 random ints.
func CreateRandomArray10000() [10000]int {
	var myArray [10000]int

	for i := 0; i < 10000; i++ {
		myArray[i] = rand.Intn(1000)
	}

	return myArray
}

// Returns a slice of the specified number of random ints.
func CreateRandomSlice(n int) []int {
	mySlice := rand.Perm(10000)

	return mySlice
}

// Returns an array with 10000 elements initialised to 1.
func createArray10000() [10000]int {
	var myArray [10000]int

	for i := 0; i < 10000; i++ {
		myArray[i] = 1
	}

	return myArray
}

// Returns a slice of n elements initialised to 1.
func CreateSlice(n int) []int {
	mySlice := rand.Perm(10000)

	for i := 0; i < n; i++ {
		mySlice[i] = 1
	}

	return mySlice
}
