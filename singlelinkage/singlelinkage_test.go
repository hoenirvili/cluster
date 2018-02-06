package singlelinkage_test

import (
	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
	"github.com/hoenirvili/cluster/singlelinkage"
	gc "gopkg.in/check.v1"
)

type singleLinkageSuite struct{}

var _ = gc.Suite(&singleLinkageSuite{})

func (s singleLinkageSuite) TestNewSingleLinkage(c *gc.C) {
	sl := singlelinkage.NewSingleLinkage()
	c.Assert(sl, gc.NotNil)
}

func (s singleLinkageSuite) TestSingleLinkageSwap(c *gc.C) {
	sl := singlelinkage.NewSingleLinkage()
	first := distance.Distance{}
	second := distance.Distance{}
	sl.Swap(first, second)
	c.Assert(first.Points, gc.IsNil)
	c.Assert(second.Points, gc.IsNil)

	first.Points = map[set.Set]float64{
		"x1": 3.11,
		"x2": 6.12,
	}

	second.Points = map[set.Set]float64{
		"x1": 0.2,
		"x3": 0.8,
		"x7": 0.1,
	}

	sl.Swap(first, second)
	c.Assert(first.Points, gc.DeepEquals, map[set.Set]float64{
		"x1": 0.2, "x2": 6.12})
}

func (s singleLinkageSuite) TestSingleLinkageRecompute(c *gc.C) {
	on := distance.Distance{
		Points: map[set.Set]float64{
			"x1": 0.2,
			"x2": 0.8,
			"x3": 0.1,
			"x7": 1.6,
		},
	}
	based := set.Set("x2,x3")
	sl := singlelinkage.NewSingleLinkage()
	best, clusters := sl.Recompute(based, on)
	c.Assert(best, gc.Equals, 0.1)
	c.Assert(clusters, gc.DeepEquals, []set.Set{"x2"})

	best, clusters = sl.Recompute(set.Set("x9,x10"), on)
	c.Assert(best, gc.Equals, -1.0)
	c.Assert(clusters, gc.DeepEquals, []set.Set{})
}
