package structs

import (
	"math"
	"testing"
)

// Table Driven Tests
// We create a series of cases and tests them within a loop.

// Tests the Perimeter of several shapes.
func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape  Shape
		expect float64
	}{
		{Rectangle{10.0, 5.0}, 30.0},
		{Rectangle{5.0, 6.0}, 22.0},
		{Circle{10.0}, 2 * math.Pi * 10.0},
		{Circle{6.0}, 2 * math.Pi * 6.0},
	}

	for _, tt := range perimeterTests {
		CheckPerimeter(t, tt.shape, tt.expect)
	}
}

// Tests the Area of several shapes.
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape  Shape
		expect float64
	}{
		{Rectangle{10.0, 5.0}, 50.0},
		{Rectangle{5.0, 6.0}, 30.0},
		{Circle{10.0}, math.Pi * 100.0},
		{Circle{6.0}, math.Pi * 36.0},
	}

	for _, tt := range areaTests {
		CheckArea(t, tt.shape, tt.expect)
	}
}

func CheckArea(t *testing.T, shape Shape, expect float64) {
	area := shape.Area()
	if area != expect {
		t.Errorf("\nresult: %f\nexpect: %f\n", area, expect)
	}
}

func CheckPerimeter(t *testing.T, shape Shape, expect float64) {
	perimeter := shape.Perimeter()
	if perimeter != expect {
		t.Errorf("\nresult: %f\nexpect: %f\n", perimeter, expect)
	}
}
