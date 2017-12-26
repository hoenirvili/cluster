package dimension

import "unsafe"

type Point interface {
	Coordinates() []float64
}

type Distancer interface {
	Point
	Distance(p Point) float64
}

type DType uint8

const (
	one DType = iota
	two
)

func NewDistancer(points unsafe.Pointer, t DType) Distancer {

}
