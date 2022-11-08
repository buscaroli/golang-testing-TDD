package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	result := Add(3, 5)
	expect := 8

	checkCorrectResult(t, result, expect)
}

func TestVariadicAdder(t *testing.T) {
	t.Run("Testing with 4 numbers", func(t *testing.T) {
		result := AddAll(2, 4, 6, 8)
		expect := 20

		checkCorrectResult(t, result, expect)
	})

	t.Run("Testing with 0 numbers", func(t *testing.T) {
		result := AddAll()
		expect := 0

		checkCorrectResult(t, result, expect)
	})

	t.Run("Testing with 1 number", func(t *testing.T) {
		result := AddAll(5)
		expect := 5

		checkCorrectResult(t, result, expect)
	})
}

func checkCorrectResult(t testing.TB, result, expect int) {
	if result != expect {
		t.Errorf("\nresult: %d\nexpect: %d", result, expect)
	}
}

func ExampleAdd() {
	result := Add(3, 5)
	fmt.Println(result)
	// Output: 8
}
func ExampleAddAll() {
	result := AddAll(1, 3, 5, 7, 9)
	fmt.Println(result)
	// Output: 25
}
