package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}

// func TestArea(t *testing.T) {
// 	AssertArea := func(t testing.TB, shape Shape, expected float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != expected {
// 			t.Errorf("expected %.2f but got %.2f", expected, got)
// 		}
// 	}
// 	t.Run("Rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{5.0, 8.0}
// 		AssertArea(t, rectangle, 40.0)
// 	})
// 	t.Run("Circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		AssertArea(t, circle, 314.1592653589793)
// 	})
// }

// Test with Table driven tests this are useful when you want to build a list of test cases
// that can be tested in the same manner.
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape    Shape
		name     string
		expected float64
	}{
		{
			name:     "Rectangle",
			shape:    Rectangle{10.0, 10.0},
			expected: 100.0,
		},
		{
			name:     "Circle",
			shape:    Circle{10},
			expected: 314.1592653589793,
		},
		{
			name:     "Triangle",
			shape:    Triangle{12, 6},
			expected: 36.0,
		},
	}

	for _, tt := range areaTests {
		// You can run specific tests within your table with
		// go test -run TestArea/[tt.name]
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expected {
				t.Errorf("%#v expected %.2f but got %.2f", tt.name, tt.expected, got)
			}
		})
	}
}
