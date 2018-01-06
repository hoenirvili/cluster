package one_test

import (
	"github.com/hoenirvili/cluster/dimension/one"
	gc "gopkg.in/check.v1"
)

type pointSuite struct{}

var _ = gc.Suite(&pointSuite{})

func (p pointSuite) TestNewPoint(c *gc.C) {
	point := one.NewPoint(-5.1)
	c.Assert(point, gc.NotNil)
	c.Assert(point, gc.DeepEquals, one.Point(-5.1))
}

func (p pointSuite) TestNewPoints(c *gc.C) {
	points := one.NewPoints(-5.1, -7.2, 1.1)
	c.Assert(points, gc.NotNil)
	n := len(points)
	c.Assert(n, gc.Equals, 3)
	c.Assert(points, gc.DeepEquals, []one.Point{-5.1, -7.2, 1.1})
}
func (p pointSuite) TestNewPointsEmpty(c *gc.C) {
	points := one.NewPoints()
	c.Assert(points, gc.IsNil)
}

func (p pointSuite) TestNewDistancesEmpty(c *gc.C) {
	distances := one.NewDistances()
	c.Assert(distances, gc.IsNil)
}

func (p pointSuite) TestNewDistances(c *gc.C) {
	distances := one.NewDistances(-5.1, -7.2, 1.1)
	c.Assert(distances, gc.NotNil)
	n := len(distances)
	c.Assert(n, gc.Equals, 3)
}

func (p pointSuite) points(c *gc.C) []one.Point {
	points := one.NewPoints(-2.2, -2.0, -0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0)
	c.Assert(points, gc.NotNil)
	c.Assert(len(points), gc.Equals, 10)
	return points
}

func (p pointSuite) TestPointCoordinates(c *gc.C) {
	points := p.points(c)
	expected := []float64{-2.2, -2.0, -0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0}
	for i, point := range points {
		got := point.Coordinates()[0]
		c.Assert(got, gc.DeepEquals, expected[i])
	}
}

func (p pointSuite) TestPointDistance(c *gc.C) {
	points := p.points(c)
	done := points[0]
	dsecond := points[1]
	got := done.Distance(&dsecond)
	c.Assert(got, gc.DeepEquals, 0.2)
}
