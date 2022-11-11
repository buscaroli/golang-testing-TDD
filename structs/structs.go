package structs

import (
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

// Calculates the Perimeter of a rectangle as a float64.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Calculates the Area oa a rectangle as a float64.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Calculates the Perimeter of a circle as a float64.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Calculates the Area of a circle as a float64.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
