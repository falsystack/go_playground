package strcut_method_interface

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func Perimeter(rec *Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}
