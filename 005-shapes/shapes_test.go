package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangle should get correct perimeter", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		perimeter := rect.Perimeter()
		expected := 40.0

		if perimeter != expected {
			t.Errorf("expected %.2f got %.2f", expected, perimeter)
		}
	})
}

func TestArea(t *testing.T) {
	tests := []struct {
		title    string
		shape    Shape
		expected float64
	}{
		{"rectangle", &Rectangle{10.0, 10.0}, 100.0},
		{"circle", &Circle{10.0}, 62.83185307179586},
		{"triangle", &Triangle{12.0, 6.0}, 36.0},
	}

	assertArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		area := shape.Area()
		if area != expected {
			t.Errorf("%#v expected %g got %g", shape, expected, area)
		}
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			assertArea(t, test.shape, test.expected)
		})
	}
}
