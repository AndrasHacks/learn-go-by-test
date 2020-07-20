package main

import "testing"

func TestStructs(t *testing.T) {
	assertEquals := func(want float64, got float64, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("Wanted %.2f, but got %.2f", want, got)
		}
	}

	t.Run("Given 10 width &  10 height rectange, When GetPerimeter(), Then perimeter returned.",
		func(t *testing.T) {
			want := 40.00
			got := GetPerimeter(Rectangle{10.00, 10.00})
			assertEquals(want, got, t)
		})

	t.Run("Given 10 width & 10 height rectange, When GetArea(), Then 100 is the area.",
		func(t *testing.T) {
			want := 100.00
			rectangle := Rectangle{10.00, 10.00}
			got := rectangle.Area()
			assertEquals(want, got, t)
		})

	t.Run("Given 10 radius circle, When GetArea(), Then 314.15 returned",
		func(t *testing.T) {
			want := 314.1592653589793
			circle := Circle{10.00}
			got := circle.Area()

			assertEquals(want, got, t)
		})
}

func TestArea(t *testing.T) {

	// In Go we have implicit interface resolution

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Wanted %.2f got %g", want, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectange := Rectangle{12, 6}
		checkArea(t, rectange, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	// Same tests as above but in a table driven mode
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Wanted %.2f got %g (%#v)", tt.want, got, tt.shape)
		}
	}
}
