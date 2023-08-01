package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("%v got %g, but want %g", shape, got, want)
		}
	}

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{6.0, 2.0}, hasArea: 12.0},
		{name: "Cicle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{5.0, 4.0}, hasArea: 10.0},
	}
	for _, s := range areaTests {
		t.Run(s.name, func(t *testing.T) {
			checkArea(t, s.shape, s.hasArea)
		})
	}
}
