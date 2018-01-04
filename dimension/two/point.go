package two

import (
	"math"

	"github.com/hoenirvili/cluster/dimension"
	"github.com/hoenirvili/cluster/util"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

var (
	_ dimension.Point     = (*Point)(nil)
	_ dimension.Distancer = (*Point)(nil)
)

func NewDistances(x, y []float64) []dimension.Distancer {
	ps := NewPoints(x, y)
	d := make([]dimension.Distancer, 0, len(ps))
	for _, point := range ps {
		d = append(d, point)
	}

	return d
}

func NewPoints(x, y []float64) []Point {
	if x == nil || y == nil {
		return nil
	}

	n := len(x)
	m := len(y)

	if n != m {
		return nil
	}

	points := make([]Point, 0, n)

	for i := 0; i < n; i++ {
		points = append(points, Point{X: x[i], Y: y[i]})
	}

	return points
}

// Coordinates creates a slice of coordinates of the
// two dimensional point
func (p Point) Coordinates() []float64 {
	return []float64{p.X, p.Y}
}

// Distance computes the distance between the fixed point
// and the given dimension.Point
// This will use euclidian distance
func (p Point) Distance(x dimension.Point) float64 {
	cord := x.Coordinates()
	px, py := cord[0], cord[1]
	p1, p2 := math.Pow((p.X-px), 2), math.Pow((p.Y-py), 2)
	psum := p1 + p2
	r := math.Sqrt(psum)
	return util.Round(r, 4)
}
