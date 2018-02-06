// Package dimension provides a small interface
// for defining different set of point
package dimension

// Point defines a given one, two or n dimension point
type Point interface {
	// Coordinates returns the coordinates of the n dimension point
	Coordinates() []float64
}

// Distancer is the minimal set of behaviour
// for computing the distance between two points
type Distancer interface {
	Point
	// Distance returns the distance
	// of the fixed point with the given point
	Distance(p Point) float64
}
