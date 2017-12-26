package cluster_test

import (
	"fmt"

	"github.com/hoenirvili/cluster"
	"github.com/hoenirvili/cluster/dimension/one"
	gc "gopkg.in/check.v1"
)

func (cs clusterSuite) TestNewDistances(c *gc.C) {
	points := one.NewPoints(-2.2, -2.0, -0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0)
	c.Assert(points, gc.NotNil)
	distances := cluster.NewDistances(points)
	c.Assert(distances, gc.NotNil)
	fmt.Println(distances)
}
