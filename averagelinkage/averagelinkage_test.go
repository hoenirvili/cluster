package averagelinkage_test

import (
	"github.com/hoenirvili/cluster/averagelinkage"
	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
	gc "gopkg.in/check.v1"
)

type averageLinkageSuite struct{}

var _ = gc.Suite(&averageLinkageSuite{})

func (a averageLinkageSuite) table() []distance.Distance {
	return []distance.Distance{
		{
			Set: "x1",
			Points: map[set.Set]float64{
				"x2": 0.32,
				"x3": 0.10,
			},
		},
		{
			Set: "x2",
			Points: map[set.Set]float64{
				"x3": 0.80,
			},
		},
		{
			Set:    "x3",
			Points: nil,
		},
	}
}

func (a averageLinkageSuite) TestNewAverageLinkage(c *gc.C) {
	avl := averagelinkage.NewAverageLinkage(nil)
	c.Assert(avl, gc.IsNil)

	table := a.table()
	avl = averagelinkage.NewAverageLinkage(table)
	c.Assert(avl, gc.NotNil)
	for i, d := range avl.Table {
		c.Assert(d.Set, gc.DeepEquals, table[i].Set)
		for key, point := range d.Points {
			c.Assert(point, gc.DeepEquals, table[i].Points[key])
		}
	}
}

func (a averageLinkageSuite) TestAverageLinkageSwap(c *gc.C) {
	table := a.table()

	avl := averagelinkage.NewAverageLinkage(table)
	var first, second distance.Distance
	avl.Swap(first, second)
	c.Assert(first.Points, gc.IsNil)
	c.Assert(second.Points, gc.IsNil)

	first, second = table[0], table[1]
	avl.Swap(first, second)

	c.Assert(first.Points, gc.DeepEquals, map[set.Set]float64{"x2": 0.32, "x3": 0.1})
	c.Assert(first.Points, gc.DeepEquals, map[set.Set]float64{"x3": 0.1, "x2": 0.32})
}

func (a averageLinkageSuite) TestAverageLinkageRecompute(c *gc.C) {
	table := a.table()

	avl := averagelinkage.NewAverageLinkage(table)
	based := set.Set("x2")
	on := distance.Distance{
		Set:    set.Set("x2"),
		Points: map[set.Set]float64{"x3": 0.80},
	}
	best, toDelete := avl.Recompute(based, on)
	// let's make sure the on it's not modified after the call
	c.Assert(on, gc.DeepEquals, distance.Distance{
		Set:    set.Set("x2"),
		Points: map[set.Set]float64{"x3": 0.80},
	})
	c.Assert(best, gc.Equals, -1.0)
	c.Assert(toDelete, gc.DeepEquals, []set.Set{})

	based = set.Set("x1,x8")
	on = distance.Distance{
		Set: set.Set("x1"),
		Points: map[set.Set]float64{
			"x2": 0.32,
			"x3": 0.10,
		},
	}
	best, toDelete = avl.Recompute(based, on)
	// let's make sure the on it's not modified after the call
	c.Assert(on, gc.DeepEquals, distance.Distance{
		Set:    set.Set("x1"),
		Points: map[set.Set]float64{"x2": 0.32, "x3": 0.10},
	})
	c.Assert(best, gc.Equals, -1.0)
	c.Assert(toDelete, gc.DeepEquals, []set.Set{})

	based = set.Set("x2,x3")
	best, toDelete = avl.Recompute(based, on)
	c.Assert(on, gc.DeepEquals, distance.Distance{
		Set:    set.Set("x1"),
		Points: map[set.Set]float64{"x3": 0.10, "x2": 0.32},
	})
	c.Assert(best, gc.Equals, 0.21)
	c.Assert(toDelete, gc.DeepEquals, []set.Set{"x2"})
}
