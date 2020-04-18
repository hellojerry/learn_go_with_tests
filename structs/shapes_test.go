package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T){
	rect := Rectangle{10.0,10.0}
	res := Perimeter(rect)
	expected := 40.0
	if res != expected {
		t.Errorf("res: %.2f, expected: %.2f",
			res,expected)
	}
}

func TestArea(t *testing.T){

	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
	{name: "Rectangle",shape:Rectangle{Width:12.0,Height:6.0},hasArea:72.0},
		{name: "Circle",shape:Circle{Radius:10.0},hasArea:314.1592653589793},
		{name: "Triangle",shape: Triangle{Base:12.0,Height:6.0},hasArea:36.0},
	}
	for _, tt := range areaTests {
		res := tt.shape.Area()
		if res != tt.hasArea {
			t.Errorf("%#v res: %.2f, expected: %.2f",
				tt.shape,res,tt.hasArea)
		}

	}

}
