package two_test

import (
	"github.com/hoenirvili/cluster/dimension/two"
	gc "gopkg.in/check.v1"
)

type pointSuite struct{}

var _ = gc.Suite(&pointSuite{})

var (
	X = []float64{-4, -3, -2, -1, 1, 1, 2, 3, 3, 4}
	Y = []float64{-2, -2, -2, -2, -1, 1, 3, 2, 4, 3}
)

func (p pointSuite) points(c *gc.C) []two.Point {
	points := two.NewPoints(X, Y)
	c.Assert(points, gc.NotNil)
	return points
}

func (p pointSuite) TestNewPoint(c *gc.C) {
	point := two.NewPoint(-1, -2)
	c.Assert(point, gc.NotNil)
	expected := two.Point{X: -1, Y: -2}
	c.Assert(point, gc.DeepEquals, expected)
}

func (p pointSuite) TestNewPoints(c *gc.C) {
	points := two.NewPoints([]float64{1, 2, 2}, []float64{2, 1, 1})
	c.Assert(points, gc.NotNil)
	n := len(points)
	c.Assert(n, gc.Equals, 3)
	c.Assert(points, gc.DeepEquals, []two.Point{{1, 2}, {2, 1}, {2, 1}})
}

func (p pointSuite) TestNewDistances(c *gc.C) {
	distances := two.NewDistances([]float64{1, 2, 2}, []float64{2, 1, 1})
	c.Assert(distances, gc.NotNil)
	n := len(distances)
	c.Assert(n, gc.Equals, 3)
}

func (p pointSuite) TestNewDistancesWithError(c *gc.C) {
	distances := two.NewDistances([]float64{1, 2, 2}, []float64{2, 1})
	c.Assert(distances, gc.IsNil)

	distances = two.NewDistances([]float64{1, 2}, []float64{2, 6, 1})
	c.Assert(distances, gc.IsNil)

	distances = two.NewDistances(nil, []float64{2, 6, 1})
	c.Assert(distances, gc.IsNil)

	distances = two.NewDistances([]float64{2, 6, 1}, nil)
	c.Assert(distances, gc.IsNil)

	distances = two.NewDistances(nil, nil)
	c.Assert(distances, gc.IsNil)
}

func (p pointSuite) TestNewPointsWithError(c *gc.C) {
	points := two.NewPoints([]float64{1, 2, 5, 2}, []float64{2, 1})
	c.Assert(points, gc.IsNil)

	points = two.NewPoints(nil, []float64{3, 1})
	c.Assert(points, gc.IsNil)

	points = two.NewPoints([]float64{3, 1}, nil)
	c.Assert(points, gc.IsNil)

	points = two.NewPoints(nil, nil)
	c.Assert(points, gc.IsNil)
}

func (p pointSuite) TestPointCoordinates(c *gc.C) {
	points := p.points(c)
	for i, point := range points {
		coordinates := point.Coordinates()
		c.Assert(len(coordinates), gc.Equals, 2)
		x := coordinates[0]
		y := coordinates[1]
		c.Assert(x, gc.Equals, X[i])
		c.Assert(y, gc.Equals, Y[i])
	}
}

func (p pointSuite) TestPointDistance(c *gc.C) {
	points := p.points(c)
	done := points[0]
	dsecond := points[1]
	got := done.Distance(&dsecond)
	c.Assert(got, gc.DeepEquals, 1.0)
}
