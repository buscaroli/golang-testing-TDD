package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Calculates perimeter of the Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		result := rectangle.Perimeter()
		expect := 30.0

		CheckCorrectResult(t, expect, result)
	})

	t.Run("Calculates circumference of Circle", func(t *testing.T) {
		circle := Circle{10.0}
		result := circle.Circumference()
		expect := 2 * math.Pi * 10.0
		CheckCorrectResult(t, result, expect)
	})
}

func TestArea(t *testing.T) {
	t.Run("Calculates the area of a Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		result := rectangle.Area()
		expect := 50.0

		CheckCorrectResult(t, result, expect)
	})

	t.Run("Calculates the area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		result := circle.Area()
		expect := math.Pi * 100.00

		CheckCorrectResult(t, result, expect)
	})
}

func CheckCorrectResult(t *testing.T, result, expect float64) {
	if result != expect {
		t.Errorf("\nresult: %f\nexpect: %f\n", result, expect)
	}
}
