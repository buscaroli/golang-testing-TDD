package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("Test Repeat() works with a string and a number larger than zero", func(t *testing.T) {
		result := Repeat("hi", 3)
		expect := "hihihi"

		checkCorrectResult(t, result, expect)
	})

	t.Run("Test Repeat() works when passing a number of repetitions equal to zero", func(t *testing.T) {
		result := Repeat("hi", 0)
		expect := ""

		checkCorrectResult(t, result, expect)
	})

	t.Run("Test RepeatWithSeparator() works when passing all arguments", func(t *testing.T) {
		result := RepeatWithSeparator("Hi", ", ", 3)
		expect := "Hi, Hi, Hi"

		checkCorrectResult(t, result, expect)
	})

	t.Run("Test RepeatWithSeparator() works when omitting the separator", func(t *testing.T) {
		result := RepeatWithSeparator("Hi", "", 4)
		expect := "HiHiHiHi"

		checkCorrectResult(t, result, expect)
	})

	t.Run("Test RepeatWithSeparator2() works when passing all arguments", func(t *testing.T) {
		result := RepeatWithSeparator2("Hi", ", ", 3)
		expect := "Hi, Hi, Hi"

		checkCorrectResult(t, result, expect)
	})

	t.Run("Test RepeatWithSeparator2() works when omitting the separator", func(t *testing.T) {
		result := RepeatWithSeparator2("Hi", "", 4)
		expect := "HiHiHiHi"

		checkCorrectResult(t, result, expect)
	})

}

func checkCorrectResult(t testing.TB, result, expect string) {
	if result != expect {
		t.Errorf("\nresult: %s\nexpect: %s\n", result, expect)
	}
}

func ExampleRepeat() {
	result := Repeat("Hello", 3)
	fmt.Println(result)
	// Output: HelloHelloHello
}

func ExampleRepeatWithSeparator() {
	result := RepeatWithSeparator("Buzz", "...", 4)
	fmt.Println(result)
	// Output : Buzz...Buzz...Buzz...Buzz
}

func BenchmarkRepeat5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("Hello", 5)
	}
}

func BenchmarkRepeatWithSeparator3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithSeparator("Hello", ", ", 5)
	}
}

func BenchmarkRepeatWithSeparatorMissingSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithSeparator("Hello", "", 5)
	}
}

func BenchmarkRepeatWithSeparator2MissingSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithSeparator2("Hello", "", 5)
	}
}
