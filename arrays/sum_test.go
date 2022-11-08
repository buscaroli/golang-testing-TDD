package arrays

import (
	"fmt"
	"testing"
)

// Constants /////////////////////////////////

var bmarkSampleArray [10000]int = CreateRandomArray10000()
var bmarkSampleSlice10000 = CreateRandomSlice(10000)

var sampleArray [10000]int = createArray10000()
var sampleSlice10000 []int = CreateSlice(10000)

// Tests //////////////////////////////////////

func testSum(t *testing.T) {
	t.Run("Checks SumIntArray returns the sum of an array of 5 ints", func(t *testing.T) {

		result := SumIntArray(sampleArray)
		expect := 10000

		checkCorrectResult(t, result, expect)
	})

	t.Run("Checks SumIntArray2 returns the sum of an array of 5 ints", func(t *testing.T) {

		result := SumIntArray2(sampleArray)
		expect := 10000

		checkCorrectResult(t, result, expect)
	})

	t.Run("Checks SumIntSlice returns the sum of an array of 5 ints", func(t *testing.T) {

		result := SumIntSlice(sampleSlice10000)
		expect := 10000

		checkCorrectResult(t, result, expect)
	})
}

func checkCorrectResult(t testing.TB, result, expect int) {
	if result != expect {
		t.Errorf("\nresult: %d\nexpect: %d\n", result, expect)
	}
}

// Examples ////////////////////////////

func ExampleSumIntArray() {

	result := SumIntArray(sampleArray)
	fmt.Println(result)
	// Output: 10000
}

func ExampleSumIntArray2() {

	result := SumIntArray2(sampleArray)
	fmt.Println(result)
	// Output: 10000
}

func ExampleSumIntSlice() {

	result := SumIntSlice(sampleSlice10000)
	fmt.Println(result)
	// Output: 10000
}

// Benchmarks //////////////////////////
// Benchmarks using the Array are twice as fast as the ones using a Slice.

func BenchmarkSumIntArrayWithFiveElems(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SumIntArray(bmarkSampleArray)
	}
}

func BenchmarkSumIntArray2WithFiveElems(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SumIntArray2(bmarkSampleArray)
	}
}

func BenchmarkSumIntSliceWithFiveElems(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SumIntSlice(bmarkSampleSlice10000)
	}
}
