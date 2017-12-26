package cluster

import (
	"fmt"

	"github.com/hoenirvili/cluster/dimension"
)

type Distance struct {
	cluster Cluster
	points  map[Cluster]float64
}

// NewDistances
func NewDistances(points []dimension.Distancer) []Distance {
	prefix := "x"
	n := len(points)
	distances := make([]Distance, 0, 0)

	for i := 0; i < n; i++ {
		// create a new fixed cluster
		distance := Distance{
			cluster: NewCluster(fmt.Sprintf("%d%d", prefix, i+1)),
		}

		// if we reached the end of the list
		// this means we are on the last distance
		// and we should skip it
		if i+1 == n {
			continue
		}

		distance.points = make(map[Cluster]float64)
		// for every left point compute the distance from the
		// fixed cluster to other clusters
		for j := i + 1; j < n; j++ {
			// length is the distance between cluster i and j
			length := points[i].Distance(points[j])
			// create the cluster name  for j
			key := NewCluster(fmt.Sprintf("%s%d", prefix, j+1))
			// assign distance to the map
			distance.points[key] = length
		}
		distances = append(distances, distance)
	}

	return distances
}
