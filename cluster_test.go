package cluster_test

import (
	"github.com/hoenirvili/cluster"
	"github.com/hoenirvili/cluster/dimension/one"
	"github.com/hoenirvili/cluster/dimension/two"
	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
	gc "gopkg.in/check.v1"
)

type clusterSuite struct{}

var _ = gc.Suite(&clusterSuite{})

func (cs clusterSuite) oneDistances(c *gc.C) []distance.Distance {
	points := one.NewDistances(-0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0)
	c.Assert(points, gc.NotNil)
	distances := distance.NewDistances(points)
	c.Assert(distances, gc.NotNil)
	return distances
}

func (cs clusterSuite) twoDistances(c *gc.C) []distance.Distance {
	points := two.NewDistances(
		[]float64{-4, -3, -2, -1, 1, 1, 2, 3, 3, 4},
		[]float64{-2, -2, -2, -2, -1, 1, 3, 2, 4, 3},
	)
	c.Assert(points, gc.NotNil)
	distances := distance.NewDistances(points)
	c.Assert(distances, gc.NotNil)
	return distances
}

func (cl clusterSuite) TestDistanceFitOneSingleLinkage(c *gc.C) {
	expected := [][]set.Set{
		{"x1,x2,x3,x4,x5,x6,x7,x8"},
		{"x1,x2,x3,x4", "x5,x6,x7,x8"},
		{"x1", "x2,x3,x4", "x5,x6,x7,x8"},
		{"x1", "x2,x3,x4", "x5,x6", "x7,x8"},
		{"x1", "x2,x3", "x4", "x5,x6", "x7,x8"},
		{"x1", "x2,x3", "x4", "x5,x6", "x7", "x8"},
		{"x1", "x2,x3", "x4", "x5", "x6", "x7", "x8"},
		{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8"},
	}

	distances := cl.oneDistances(c)
	iter := len(distances)
	for i := iter; i > 0; i-- {
		clusters := cluster.Fit(distances, cluster.SingleLinkage, i)
		c.Assert(clusters, gc.NotNil)
		n := len(clusters)
		c.Assert(n, gc.Equals, i)
		c.Assert(clusters, gc.DeepEquals, expected[i-1])
	}
}

func (cl clusterSuite) TestDistanceFitOneCompleteLinkage(c *gc.C) {
	expected := [][]set.Set{
		{"x1,x2,x3,x4,x5,x6,x7,x8"},
		{"x1,x2,x3,x4", "x5,x6,x7,x8"},
		{"x1", "x2,x3,x4", "x5,x6,x7,x8"},
		{"x1", "x2,x3,x4", "x5,x6", "x7,x8"},
		{"x1", "x2,x3", "x4", "x5,x6", "x7,x8"},
		{"x1", "x2,x3", "x4", "x5,x6", "x7", "x8"},
		{"x1", "x2,x3", "x4", "x5", "x6", "x7", "x8"},
		{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8"},
	}

	distances := cl.oneDistances(c)
	iter := len(distances)
	for i := iter; i > 0; i-- {
		clusters := cluster.Fit(distances, cluster.CompleteLinkage, i)
		c.Assert(clusters, gc.NotNil)
		n := len(clusters)
		c.Assert(n, gc.Equals, i)
		c.Assert(clusters, gc.DeepEquals, expected[i-1])
	}
}

func (cl clusterSuite) TestDistanceFitOneAverageLinkage(c *gc.C) {
	expected := [][]set.Set{
		{"x1,x2,x3,x4,x5,x6,x7,x8"},
		{"x1,x2,x3,x4", "x5,x6,x7,x8"},
		{"x1", "x2,x3,x4", "x5,x6,x7,x8"},
		{"x1", "x2,x3,x4", "x5,x6", "x7,x8"},
		{"x1", "x2,x3", "x4", "x5,x6", "x7,x8"},
		{"x1", "x2,x3", "x4", "x5,x6", "x7", "x8"},
		{"x1", "x2,x3", "x4", "x5", "x6", "x7", "x8"},
		{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8"},
	}

	distances := cl.oneDistances(c)
	iter := len(distances)
	for i := iter; i > 0; i-- {
		clusters := cluster.Fit(distances, cluster.AverageLinkage, i)
		c.Assert(clusters, gc.NotNil)
		n := len(clusters)
		c.Assert(n, gc.Equals, i)
		c.Assert(clusters, gc.DeepEquals, expected[i-1])
	}
}

func (cl clusterSuite) TestDistanceFitTwoSingleLinkage(c *gc.C) {
	expected := [][]set.Set{
		{"x1,x2,x3,x4,x5,x6,x7,x8,x9,x10"},
		{"x1,x2,x3,x4,x5,x6", "x7,x8,x9,x10"},
		{"x1,x2,x3,x4", "x5,x6", "x7,x8,x9,x10"},
		{"x1,x2,x3,x4", "x5", "x6", "x7,x8,x9,x10"},
		{"x1,x2,x3,x4", "x5", "x6", "x7,x8,x9", "x10"},
		{"x1,x2,x3,x4", "x5", "x6", "x7,x8", "x9", "x10"},
		{"x1,x2,x3,x4", "x5", "x6", "x7", "x8", "x9", "x10"},
		{"x1,x2,x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10"},
		{"x1,x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10"},
		{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10"},
	}

	distances := cl.twoDistances(c)
	iter := len(distances)
	for i := iter; i > 0; i-- {
		clusters := cluster.Fit(distances, cluster.SingleLinkage, i)
		c.Assert(clusters, gc.NotNil)
		n := len(clusters)
		c.Assert(n, gc.Equals, i)
		c.Assert(clusters, gc.DeepEquals, expected[i-1])
	}
}

func (cl clusterSuite) TestDistanceFitTwoCompleteLinkage(c *gc.C) {
	expected := [][]set.Set{
		{"x1,x2,x3,x4,x5,x6,x7,x8,x9,x10"},
		{"x1,x2,x3,x4", "x5,x6,x7,x8,x9,x10"},
		{"x1,x2,x3,x4", "x5,x6", "x7,x8,x9,x10"},
		{"x1,x2", "x3,x4", "x5,x6", "x7,x8,x9,x10"},
		{"x1,x2", "x3,x4", "x5,x6", "x7,x8", "x9,x10"},
		{"x1,x2", "x3,x4", "x5", "x6", "x7,x8", "x9,x10"},
		{"x1,x2", "x3,x4", "x5", "x6", "x7,x8", "x9", "x10"},
		{"x1,x2", "x3,x4", "x5", "x6", "x7", "x8", "x9", "x10"},
		{"x1,x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10"},
		{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10"},
	}

	distances := cl.twoDistances(c)
	iter := len(distances)
	for i := iter; i > 0; i-- {
		clusters := cluster.Fit(distances, cluster.CompleteLinkage, i)
		c.Assert(clusters, gc.NotNil)
		n := len(clusters)
		c.Assert(n, gc.Equals, i)
		c.Assert(clusters, gc.DeepEquals, expected[i-1])
	}
}

func (cl clusterSuite) TestDistanceFitTwoAverageLinkage(c *gc.C) {
	expected := [][]set.Set{
		{"x1,x2,x3,x4,x5,x6,x7,x8,x9,x10"},
		{"x1,x2,x3,x4", "x5,x6,x7,x8,x9,x10"},
		{"x1,x2,x3,x4", "x5,x6", "x7,x8,x9,x10"},
		{"x1,x2,x3,x4", "x5", "x6", "x7,x8,x9,x10"},
		{"x1,x2", "x3,x4", "x5", "x6", "x7,x8,x9,x10"},
		{"x1,x2", "x3,x4", "x5", "x6", "x7,x8", "x9,x10"},
		{"x1,x2", "x3,x4", "x5", "x6", "x7,x8", "x9", "x10"},
		{"x1,x2", "x3,x4", "x5", "x6", "x7", "x8", "x9", "x10"},
		{"x1,x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10"},
		{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10"},
	}

	distances := cl.twoDistances(c)
	iter := len(distances)
	for i := iter; i > 0; i-- {
		clusters := cluster.Fit(distances, cluster.AverageLinkage, i)
		c.Assert(clusters, gc.NotNil)
		n := len(clusters)
		c.Assert(n, gc.Equals, i)
		c.Assert(clusters, gc.DeepEquals, expected[i-1])
	}
}

// benchmarks

func (cl clusterSuite) BenchmarkFitOneSingleLinkage(c *gc.C) {
	distances := cl.oneDistances(c)
	for i := 0; i < c.N; i++ {
		iter := len(distances)
		for i := iter; i > 0; i-- {
			cluster.Fit(distances, cluster.SingleLinkage, i)
		}
	}
}

func (cl clusterSuite) BenchmarkFitOneCompleteLinkage(c *gc.C) {
	distances := cl.oneDistances(c)
	for i := 0; i < c.N; i++ {
		iter := len(distances)
		for i := iter; i > 0; i-- {
			cluster.Fit(distances, cluster.CompleteLinkage, i)
		}
	}
}

func (cl clusterSuite) BenchmarkFitOneAverageLinkage(c *gc.C) {
	distances := cl.oneDistances(c)
	for i := 0; i < c.N; i++ {
		iter := len(distances)
		for i := iter; i > 0; i-- {
			cluster.Fit(distances, cluster.AverageLinkage, i)
		}
	}
}

func (cl clusterSuite) BenchmarkFitTwoSingleLinkage(c *gc.C) {
	distances := cl.twoDistances(c)
	for i := 0; i < c.N; i++ {
		iter := len(distances)
		for i := iter; i > 0; i-- {
			cluster.Fit(distances, cluster.SingleLinkage, i)
		}
	}
}

func (cl clusterSuite) BenchmarkFitTwoCompleteLinkage(c *gc.C) {
	distances := cl.twoDistances(c)
	for i := 0; i < c.N; i++ {
		iter := len(distances)
		for i := iter; i > 0; i-- {
			cluster.Fit(distances, cluster.CompleteLinkage, i)
		}
	}
}

func (cl clusterSuite) BenchmarkFitTwoAverageLinkage(c *gc.C) {
	distances := cl.twoDistances(c)
	for i := 0; i < c.N; i++ {
		iter := len(distances)
		for i := iter; i > 0; i-- {
			cluster.Fit(distances, cluster.AverageLinkage, i)
		}
	}
}
