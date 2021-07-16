package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f hasArea %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	assertArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g hasArea %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		assertArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		assertArea(t, circle, want)
	})
}

// TestAreaUsingTableDrivenTests - same intent as function above
func TestAreaUsingTableDrivenTests(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	} {
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.hasArea {
				t.Errorf("%#v got %g hasArea %g", testCase.shape, got, testCase.hasArea)
			}
		})
	}
}