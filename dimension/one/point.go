// Package one defines one dimensional points
package one

import (
	"math"

	"github.com/hoenirvili/cluster/dimension"
	"github.com/hoenirvili/cluster/util"
)

// Point in one dimensions
type Point float64

// NewPoint creates a one dimension new point
func NewPoint(point float64) Point {
	return Point(point)
}

// NewPoints returns a slice of one dimension points
func NewPoints(points ...float64) []Point {
	n := len(points)
	if n == 0 {
		return nil
	}
	p := make([]Point, 0, n)
	for i := 0; i < n; i++ {
		p = append(p, Point(points[i]))
	}

	return p
}

// NewDistances returns new one dimension distance points
func NewDistances(points ...float64) []dimension.Distancer {
	ps := NewPoints(points...)
	if ps == nil {
		return nil
	}
	d := make([]dimension.Distancer, 0, len(points))
	for _, point := range ps {
		d = append(d, point)
	}
	return d
}

var (
	_ dimension.Point     = (*Point)(nil)
	_ dimension.Distancer = (*Point)(nil)
)

// Coordinates creates a slice of coordinates of the
// one dimensional point
func (p Point) Coordinates() []float64 {
	return []float64{float64(p)}
}

// Distance computes the distance between the fixed point
// and the given dimension.Point
func (p Point) Distance(x dimension.Point) float64 {
	fp := float64(p)
	fx := x.Coordinates()[0]

	if fp == 0.0 || fx == 0.0 {
		return fp + fx
	}

	if fp < 0.0 && fx < 0.0 || fx > 0.0 && fp > 0.0 {
		fx = math.Abs(fx)
		fp = math.Abs(fp)
		if fx > fp {
			return util.Round(fx-fp, 4)
		}
		return util.Round(fp-fx, 4)
	}

	if fx < 0.0 && fp > 0.0 || fx > 0.0 && fp < 0.0 {
		fx = math.Abs(fx)
		fp = math.Abs(fp)
	}

	return util.Round(fp+fx, 4)
}
