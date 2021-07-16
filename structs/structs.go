package structs

func Perimeter(rectangle Rectangle) float64 {
	return 2*(rectangle.Length + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Length * rectangle.Width
}

type Rectangle struct {
	Length float64
	Width float64
}