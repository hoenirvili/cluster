package cluster_test

import (
	"github.com/hoenirvili/cluster/dimension/one"
	"github.com/hoenirvili/cluster/dimension/two"
	"github.com/hoenirvili/cluster/distance"
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

// func (cs clusterSuite) TestNewDistances(c *gc.C) {
// 	expected := []cluster.Distance{
// 		{
// 			Cluster: "x1",
// 			Points: map[cluster.Cluster]float64{
// 				"x6": 2.00,
// 				"x7": 2.20,
// 				"x8": 2.30,
// 				"x2": 0.40,
// 				"x3": 0.50,
// 				"x4": 0.70,
// 				"x5": 1.90,
// 			},
// 		},
// 		{
// 			Cluster: "x2",
// 			Points: map[cluster.Cluster]float64{
// 				"x3": 0.10,
// 				"x4": 0.30,
// 				"x5": 1.50,
// 				"x6": 1.60,
// 				"x7": 1.80,
// 				"x8": 1.90,
// 			},
// 		},
// 		{
// 			Cluster: "x3",
// 			Points: map[cluster.Cluster]float64{
// 				"x5": 1.40,
// 				"x6": 1.50,
// 				"x7": 1.70,
// 				"x8": 1.80,
// 				"x4": 0.20,
// 			},
// 		},
// 		{
// 			Cluster: "x4",
// 			Points: map[cluster.Cluster]float64{
// 				"x5": 1.20,
// 				"x6": 1.30,
// 				"x7": 1.50,
// 				"x8": 1.60,
// 			},
// 		},
// 		{
// 			Cluster: "x5",
// 			Points: map[cluster.Cluster]float64{
// 				"x8": 0.40,
// 				"x6": 0.10,
// 				"x7": 0.30,
// 			},
// 		},
// 		{
// 			Cluster: "x6",
// 			Points: map[cluster.Cluster]float64{
// 				"x7": 0.20,
// 				"x8": 0.30,
// 			},
// 		},
// 		{
// 			Cluster: "x7",
// 			Points: map[cluster.Cluster]float64{
// 				"x8": 0.10,
// 			},
// 		},
// 	}
// 	points := one.NewDistances(-0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0)
// 	c.Assert(points, gc.NotNil)
// 	distances := cluster.NewDistances(points)
// 	c.Assert(distances, gc.NotNil)
// 	for index, distance := range distances {
// 		one := distance
// 		two := expected[index]
// 		c.Assert(one.Cluster, gc.DeepEquals, two.Cluster)
// 		for index, p := range one.Points {
// 			c.Assert(p, gc.DeepEquals, two.Points[index])
// 		}
// 	}
// }

// func (cl clusterSuite) TestDistanceBest(c *gc.C) {
// 	distances := cl.distances(c)
// 	pair := [][]cluster.Cluster{
// 		{"x1", "x2"},
// 		{"x2", "x3"},
// 		{"x3", "x4"},
// 		{"x4", "x5"},
// 		{"x5", "x6"},
// 		{"x6", "x7"},
// 		{"x7", "x8"},
// 	}

// 	bestDistances := []float64{0.40, 0.10, 0.20, 1.20, 0.10, 0.20, 0.10}

// 	for i, distance := range distances {
// 		first, second, d := distance.Best()
// 		clusters := pair[i]
// 		c.Assert(first, gc.DeepEquals, distance.Cluster)
// 		c.Assert(clusters[0], gc.Equals, first)
// 		c.Assert(clusters[1], gc.Equals, second)
// 		c.Assert(d, gc.Equals, bestDistances[i])
// 	}
// }

// func (cl clusterSuite) TestDistanceMerge(c *gc.C) {
// 	distance := cl.distances(c)[0]

// 	first, second, _ := distance.Best()
// 	c.Assert(first, gc.Equals, distance.Cluster)
// 	distance.Merge(second)
// 	first.Add(second)
// 	c.Assert(distance.Cluster, gc.DeepEquals, first)
// 	_, ok := distance.Points[second]
// 	c.Assert(ok, gc.Equals, false)
// }

// func (cl clusterSuite) TestDistanceRefit(c *gc.C) {
// 	expected := []cluster.Distance{
// 		{
// 			Cluster: "x1,x2",
// 			Points: map[cluster.Cluster]float64{
// 				"x6": 1.60,
// 				"x7": 1.80,
// 				"x8": 1.90,
// 				"x3": 0.10,
// 				"x4": 0.30,
// 				"x5": 1.50,
// 			},
// 		},
// 		{
// 			Cluster: "x3",
// 			Points: map[cluster.Cluster]float64{
// 				"x5": 1.40,
// 				"x6": 1.50,
// 				"x7": 1.70,
// 				"x8": 1.80,
// 				"x4": 0.20,
// 			},
// 		},
// 		{
// 			Cluster: "x4",
// 			Points: map[cluster.Cluster]float64{
// 				"x5": 1.20,
// 				"x6": 1.30,
// 				"x7": 1.50,
// 				"x8": 1.60,
// 			},
// 		},
// 		{
// 			Cluster: "x5",
// 			Points: map[cluster.Cluster]float64{
// 				"x8": 0.40,
// 				"x6": 0.10,
// 				"x7": 0.30,
// 			},
// 		},
// 		{
// 			Cluster: "x6",
// 			Points: map[cluster.Cluster]float64{
// 				"x7": 0.20,
// 				"x8": 0.30,
// 			},
// 		},
// 		{
// 			Cluster: "x7",
// 			Points: map[cluster.Cluster]float64{
// 				"x8": 0.10,
// 			},
// 		},
// 	}

// 	distances := cl.distances(c)

// 	first, second, _ := distances[0].Best()
// 	distances[0].Merge(second)
// 	first.Add(second)
// 	cluster.Refit(&distances, first, second, cluster.SingleLinkage)
// 	for i, d := range distances {
// 		one := d
// 		second := expected[i]
// 		c.Assert(one, gc.DeepEquals, second)
// 	}

// }

// func (cl clusterSuite) TestDistanceRefitIter(c *gc.C) {
// 	distances := cl.distances(c)
// 	n := len(distances)
// 	min := -1.0
// 	var (
// 		min_cluster cluster.Cluster
// 		s           cluster.Cluster
// 		j           int
// 	)
// 	for i := 0; i < n; i++ {
// 		first, second, distance := distances[i].Best()
// 		if min == -1.0 || min > distance {
// 			min = distance
// 			first.Add(second)
// 			min_cluster = first
// 			s = second
// 			j = i
// 		}
// 	}

// 	distances[j].Merge(s)
// 	fmt.Println(min_cluster)
// 	cluster.Refit(&distances, min_cluster, s, cluster.SingleLinkage)

// }

// func (cl clusterSuite) TestDistanceFit(c *gc.C) {
// 	distances := cl.distances(c)
// 	for _, d := range distances {
// 		fmt.Println(d)

// 	}

// 	// -0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0
// 	clusters := cluster.Fit(distances, cluster.SingleLinkage, 3)
// 	fmt.Println(clusters)
// }

// func (cl clusterSuite) TestDistanceFitComplete(c *gc.C) {
// 	distances := cl.distances(c)
// 	for _, d := range distances {
// 		fmt.Println(d)

// 	}

// 	//-0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0
// 	clusters := cluster.Fit(distances, cluster.CompleteLinkage, 2)
// 	fmt.Println(clusters)
// }

// func (cl clusterSuite) TestCustomDistanceFit(c *gc.C) {
// 	distances := []distance.Distance{
// 		{
// 			Set: set.Set("x1"),
// 			Points: map[set.Set]float64{
// 				"x2": 0.12,
// 				"x3": 0.51,
// 				"x4": 0.84,
// 				"x5": 0.28,
// 				"x6": 0.34,
// 			},
// 		},
// 		{
// 			Set: set.Set("x2"),
// 			Points: map[set.Set]float64{
// 				"x3": 0.25,
// 				"x4": 0.16,
// 				"x5": 0.77,
// 				"x6": 0.61,
// 			},
// 		},
// 		{
// 			Set: set.Set("x3"),
// 			Points: map[set.Set]float64{
// 				"x4": 0.14,
// 				"x5": 0.70,
// 				"x6": 0.93,
// 			},
// 		},
// 		{
// 			Set: set.Set("x4"),
// 			Points: map[set.Set]float64{
// 				"x5": 0.75,
// 				"x6": 0.20,
// 			},
// 		},
// 		{
// 			Set: set.Set("x5"),
// 			Points: map[set.Set]float64{
// 				"x6": 0.67,
// 			},
// 		},
// 		{
// 			Set:    set.Set("x6"),
// 			Points: nil,
// 		},
// 	}

// 	for _, d := range distances {
// 		fmt.Println(d)
// 	}

// 	clusters := cluster.Fit(distances, cluster.AverageLinkage, 2)
// 	fmt.Println(clusters)
// }

// func (cl clusterSuite) TestCustomDistanceFit(c *gc.C) {
// 	distances := cl.twoDistances(c)
// 	clusters := cluster.Fit(distances, cluster.AverageLinkage, 3)
// 	fmt.Println(clusters)
// }
