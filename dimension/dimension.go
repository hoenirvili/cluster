package dimension


type Point interface {
	Coordinates() []float64
}

type Distancer interface {
	Point
	Distance(p Point) float64
}