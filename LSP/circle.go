package main

import "math"

// Circle represents a circle shape

type Circle struct {
	Radius float64
}

// Area calculation for the circle

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
