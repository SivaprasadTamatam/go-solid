package main

// Rectangle represents a rectangle shape.
type Rectangle struct {
	Width  float64
	Height float64
}

// Calculate area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
