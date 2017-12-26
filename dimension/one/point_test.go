package one_test

import (
	"github.com/hoenirvili/cluster/dimension/one"
	gc "gopkg.in/check.v1"
)

type pointSuite struct{}

var _ = gc.Suite(&pointSuite{})

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
