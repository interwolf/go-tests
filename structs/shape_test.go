package structs

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rect{3.0, 4.0}
	got := r.Perimeter()
	want := 14.0
	if got != want {
		t.Errorf("got %.g want %.g", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rect", shape: Rect{3, 4}, area: 12.0},
		{name: "Circle", shape: Circle{1}, area: 3.14},
		{name: "Triangle", shape: Triagle{5, 6}, area: 15.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.area {
				t.Errorf("got %g, want %g", got, tt.area)
			}
		})
	}
}
