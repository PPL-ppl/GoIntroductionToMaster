package main

import "testing"

func Test_struct_func_interface(t *testing.T) {
	t.Run("周长", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("面积", func(t *testing.T) {
		got := Area(12.0, 6.0)
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("周长-Struct", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := PerimeterStruct(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("面积-Struct", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := AreaStruct(rectangle)
		want := 100.00
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("长方形面积-func", func(t *testing.T) {
		rectangle := Rectangle{20, 5}
		got := rectangle.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("圆面积-func", func(t *testing.T) {
		circle := Circle{20}
		got := circle.Area()
		want := 400.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("给定多个几何图像，符合条件求面积，不符合报错", func(t *testing.T) {
		checkArea := func(t *testing.T, shape Shape, want float64) {
			t.Helper()
			got := shape.Area()
			if got != want {
				t.Errorf("got %.2f want %.2f", got, want)
			}
		}
		t.Run("rectangles", func(t *testing.T) {
			rectangle := Rectangle{12, 6}
			checkArea(t, rectangle, 72.0)
		})

		t.Run("circles", func(t *testing.T) {
			circle := Circle{10}
			checkArea(t, circle, 100)
		})
	})
	t.Run("表格驱动测试", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 100},
			{Triangle{2, 3}, 3},
		}
		for _, tt := range areaTests {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		}
	})

}

func Area(width float64, length float64) float64 {
	return width * length
}

func Perimeter(width float64, length float64) float64 {
	return 2 * (width + length)
}

// 长方形
type Rectangle struct {
	//宽
	Width float64
	//长
	Length float64
}

func PerimeterStruct(rectangle Rectangle) interface{} {
	return 2 * (rectangle.Width + rectangle.Length)
}

func AreaStruct(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Length
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

// 圆
type Circle struct {
	//半径
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
