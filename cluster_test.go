package cluster_test

import (
	"github.com/hoenirvili/cluster"
	gc "gopkg.in/check.v1"
)

type clusterSuite struct{}

var _ = gc.Suite(&clusterSuite{})

func (cs clusterSuite) new(c *gc.C) cluster.Cluster {
	cls := cluster.NewCluster("x2", "x3", "x1")
	c.Assert(cls, gc.NotNil)
	c.Assert(cls, gc.DeepEquals, cluster.Cluster("x1,x2,x3"))
	return cls
}

func (cs clusterSuite) TestClusterSlice(c *gc.C) {
	cluster := cs.new(c)
	slice := cluster.Slice()
	c.Assert(slice, gc.DeepEquals, []string{"x1", "x2", "x3"})
}

func (cs clusterSuite) TestClusterString(c *gc.C) {
	expected := "{ x1,x2,x3 }"
	cluster := cs.new(c)
	got := cluster.String()
	c.Assert(got, gc.DeepEquals, expected)
}

func (cs clusterSuite) TestClusterAdd(c *gc.C) {
	cls := cs.new(c)
	point := cluster.Cluster("x5")
	cls.Add(point)
	expected := cluster.Cluster("x1,x2,x3," + point)
	c.Assert(cls, gc.DeepEquals, expected)
}

func (cs clusterSuite) TestClusterAddDupplicate(c *gc.C) {
	cls := cs.new(c)
	point := cluster.Cluster("x5")
	cls.Add(point)
	expected := cluster.Cluster("x1,x2,x3," + point)
	c.Assert(cls, gc.DeepEquals, expected)
	cls.Add(point)
	c.Assert(cls, gc.DeepEquals, expected)
}

func (cs clusterSuite) TestClusterDelete(c *gc.C) {
	cls := cs.new(c)
	point := cluster.Cluster("x3")
	cls.Delete(point)
	expected := cluster.Cluster("x1,x2")
	c.Assert(cls, gc.DeepEquals, expected)
}

func (cs clusterSuite) TestClusterLen(c *gc.C) {
	cls := cs.new(c)
	got := cls.Len()
	expected := 3
	c.Assert(got, gc.Equals, expected)
}

func (cs clusterSuite) TestClusterSwap(c *gc.C) {
	cls := cs.new(c)
	cls.Swap(0, 2)
	slice := cls.Slice()
	c.Assert(slice, gc.DeepEquals, []string{"x3", "x2", "x1"})
}

func (cs clusterSuite) TestClusterLess(c *gc.C) {
	cls := cs.new(c)
	got := cls.Less(0, 2)
	c.Assert(got, gc.Equals, true)
}
