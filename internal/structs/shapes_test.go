package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		in    string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12, 6}, 72.0},
		{"circle", Circle{10}, 314.1592653589793},
		{"triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
