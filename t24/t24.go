package t24

import (
	"fmt"
	"math"
	"unsafe"
)

type Point struct {
	// структура лежит в памяти последовательно [&P][&P+1]...[&P+8]...
	x, y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1 *Point, p2 *Point) float64 {
	p1_x := *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p1))))
	p1_y := *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + 8))

	p2_x := *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p2))))
	p2_y := *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p2)) + 8))

	return math.Sqrt((p1_x-p2_x)*(p1_x-p2_x) + (p1_y-p2_y)*(p1_y-p2_y))
}

func DoTask() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(5, 12)

	fmt.Println(Distance(p1, p2))
}
