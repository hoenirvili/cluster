package distance_test

import (
	"fmt"

	"github.com/hoenirvili/cluster/dimension/one"
	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
	gc "gopkg.in/check.v1"
)

type distanceSuite struct{}

var _ = gc.Suite(&distanceSuite{})

func (d distanceSuite) distances(c *gc.C) []distance.Distance {
	points := one.NewDistances(-0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0)
	c.Assert(points, gc.NotNil)
	distances := distance.NewDistances(points)
	c.Assert(distances, gc.NotNil)
	return distances
}

func (d distanceSuite) TestDistanceString(c *gc.C) {
	distance := distance.Distance{
		Set: "x1",
		Points: map[set.Set]float64{
			"x6": 2.00,
		},
	}

	str := fmt.Sprintf("%s", distance)
	c.Assert(str, gc.Equals, "{x1} => {x6}:2.00 ")
}

func (d distanceSuite) hardCodedOneDistances() []distance.Distance {
	return []distance.Distance{
		{
			Set: "x1",
			Points: map[set.Set]float64{
				"x6": 2.00,
				"x7": 2.20,
				"x8": 2.30,
				"x2": 0.40,
				"x3": 0.50,
				"x4": 0.70,
				"x5": 1.90,
			},
		},
		{
			Set: "x2",
			Points: map[set.Set]float64{
				"x3": 0.10,
				"x4": 0.30,
				"x5": 1.50,
				"x6": 1.60,
				"x7": 1.80,
				"x8": 1.90,
			},
		},
		{
			Set: "x3",
			Points: map[set.Set]float64{
				"x5": 1.40,
				"x6": 1.50,
				"x7": 1.70,
				"x8": 1.80,
				"x4": 0.20,
			},
		},
		{
			Set: "x4",
			Points: map[set.Set]float64{
				"x5": 1.20,
				"x6": 1.30,
				"x7": 1.50,
				"x8": 1.60,
			},
		},
		{
			Set: "x5",
			Points: map[set.Set]float64{
				"x8": 0.40,
				"x6": 0.10,
				"x7": 0.30,
			},
		},
		{
			Set: "x6",
			Points: map[set.Set]float64{
				"x7": 0.20,
				"x8": 0.30,
			},
		},
		{
			Set: "x7",
			Points: map[set.Set]float64{
				"x8": 0.10,
			},
		},
		{
			Set:    "x8",
			Points: nil,
		},
	}
}

func (d distanceSuite) TestNewDistances(c *gc.C) {
	expected := d.hardCodedOneDistances()
	points := one.NewDistances(-0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0)
	c.Assert(points, gc.NotNil)
	distances := distance.NewDistances(points)
	c.Assert(distances, gc.NotNil)
	for index, distance := range distances {
		one := distance
		two := expected[index]
		c.Assert(one.Set, gc.DeepEquals, two.Set)
		c.Assert(one.Points, gc.DeepEquals, two.Points)
	}
}

func (d distanceSuite) TestDistanceBest(c *gc.C) {
	distances := d.distances(c)
	pair := [][]set.Set{
		{"x1,x2", "x2"},
		{"x2,x3", "x3"},
		{"x3,x4", "x4"},
		{"x4,x5", "x5"},
		{"x5,x6", "x6"},
		{"x6,x7", "x7"},
		{"x7,x8", "x8"},
		{"", ""},
	}

	bestDistances := []float64{0.40, 0.10, 0.20, 1.20, 0.10, 0.20, 0.10, 0.0}
	// var first, second set.Set
	for i, distance := range distances {
		first, second, d := distance.Best()
		firstExpected, secondExpected := pair[i][0], pair[i][1]
		c.Assert(first, gc.Equals, firstExpected)
		c.Assert(second, gc.Equals, secondExpected)
		distanceExpected := bestDistances[i]
		c.Assert(d, gc.Equals, distanceExpected)

	}
}

func (d distanceSuite) TestDistanceMerge(c *gc.C) {
	distance := d.distances(c)[0]

	first, second, _ := distance.Best()
	distance.Merge(second)
	c.Assert(first, gc.DeepEquals, distance.Set)
	_, ok := distance.Points[second]
	c.Assert(ok, gc.Equals, false)
}

func (d distanceSuite) TestDistanceRefit(c *gc.C) {
	expected := d.hardCodedOneDistances()
	expected[0].Set = set.Set("x1,x2")
	delete(expected[0].Points, set.Set("x2"))
	distances := d.distances(c)
	first, second, _ := distances[0].Best()
	distances[0].Merge(second)
	c.Assert(distances[0].Set, gc.Equals, first)
	c.Assert(distances, gc.DeepEquals, expected)
}
