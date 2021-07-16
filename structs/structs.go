package structs

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2*(rectangle.Length + rectangle.Width)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Length * rectangle.Width
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base)/2
}

type Rectangle struct {
	// A struct is just a named collection of fields where you can store data.
	Length float64
	Width float64
}
type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}