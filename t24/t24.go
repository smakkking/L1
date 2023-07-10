package t24

import (
	"math"
	"time"
)

type Point struct {
	x, y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)

	time.Sleep()
}
