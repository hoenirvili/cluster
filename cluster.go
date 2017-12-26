// Package cluster provides a sorted set to represent clusters
// from the given points
package cluster

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Cluster describes a set of points with the naming convention
// of x1,x2,x3,x4,x5, x6x7x8, x9x10
// These will be always sorted based on their suffix order
type Cluster string

var (
	_ fmt.Stringer   = (*Cluster)(nil)
	_ sort.Interface = (*Cluster)(nil)
)

// Create a new cluster from the given points
func NewCluster(points ...string) Cluster {
	if len(points) == 1 {
		return Cluster(points[0])
	}

	cluster := Cluster(strings.Join(points, ","))
	sort.Sort(&cluster)
	return cluster
}

func (c Cluster) Slice() []string {
	return strings.Split(string(c), ",")
}

func (c Cluster) num(i int) int {
	slice := c.Slice()
	suffix := slice[i][1:]
	num, err := strconv.ParseInt(suffix, 10, 64)
	if err != nil {
		sep := strings.Split(suffix, "x")
		num, err := strconv.ParseInt(sep[0], 10, 64)
		if err != nil {
			panic(err)
		}
		return int(num)
	}

	return int(num)
}

func (c Cluster) Len() int { return len(c.Slice()) }
func (c *Cluster) Swap(i, j int) {
	slice := c.Slice()
	slice[i], slice[j] = slice[j], slice[i]
	*c = Cluster(strings.Join(slice, ","))
}
func (c Cluster) Less(i, j int) bool { return c.num(i) < c.num(j) }

// Add appends a point in the cluster
// If the point is already in the cluster it will be no-op
func (c *Cluster) Add(point string) {
	slice := c.Slice()
	n := len(slice)

	for i := 0; i < n; i++ {
		if slice[i] == point {
			return
		}
	}

	slice = append(slice, point)
	*c = Cluster(strings.Join(slice, ","))
	sort.Sort(c)
}

// Delete a given point from the cluster
func (c *Cluster) Delete(point string) {
	slice := c.Slice()
	n := len(slice)
	if n == 0 {
		return
	}

	i := -1
	for j := 0; j < n; j++ {
		if slice[j] == point {
			i = j
			break
		}
	}

	if i == -1 {
		return
	}

	slice = append(slice[:i], slice[i+1:]...)
	*c = Cluster(strings.Join(slice, ","))
	sort.Sort(c)
}

// String returns the cluster representation as a string
func (c *Cluster) String() string {
	return fmt.Sprintf("{ %s }", *c)
}
