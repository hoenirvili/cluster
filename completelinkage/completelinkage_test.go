package completelinkage_test

import (
	"github.com/hoenirvili/cluster/completelinkage"
	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
	gc "gopkg.in/check.v1"
)

type completeLinkageSuite struct{}

var _ = gc.Suite(&completeLinkageSuite{})

func (cls completeLinkageSuite) TestNewCompleteLinkage(c *gc.C) {
	cl := completelinkage.NewCompleteLinkage()
	c.Assert(cl, gc.NotNil)
}

func (cls completeLinkageSuite) TestCompleteLinkageSwap(c *gc.C) {
	cl := completelinkage.NewCompleteLinkage()
	first := distance.Distance{}
	second := distance.Distance{}
	cl.Swap(first, second)
	c.Assert(first.Points, gc.IsNil)
	c.Assert(second.Points, gc.IsNil)

	first.Points = map[set.Set]float64{"x1": 0.7, "x2": 0.1}
	second.Points = map[set.Set]float64{"x1": 1.8, "x2": 0.01}

	cl.Swap(first, second)
	c.Assert(first.Points, gc.DeepEquals, map[set.Set]float64{
		"x1": 1.8, "x2": 0.1,
	})
	c.Assert(second.Points, gc.DeepEquals, map[set.Set]float64{
		"x1": 1.8, "x2": 0.01,
	})
}

func (cls completeLinkageSuite) TestCompleteLinkageRecompute(c *gc.C) {
	based := set.Set("x2,x3")
	on := distance.Distance{
		Points: map[set.Set]float64{
			"x1": 0.2,
			"x2": 0.8,
			"x3": 0.1,
			"x7": 1.6,
		},
	}
	cl := completelinkage.NewCompleteLinkage()

	best, clusters := cl.Recompute(based, on)
	c.Assert(best, gc.Equals, 0.8)
	c.Assert(clusters, gc.DeepEquals, []set.Set{"x3"})

	best, clusters = cl.Recompute(set.Set("x9,x10"), on)
	c.Assert(best, gc.Equals, -1.0)
	c.Assert(clusters, gc.DeepEquals, []set.Set{})
}
