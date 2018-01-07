// Package used for computing distance between points
package distance

import (
	"fmt"
	"math"

	"github.com/hoenirvili/cluster/dimension"
	"github.com/hoenirvili/cluster/set"
)

// NewDistances returns a table of cluster distances
func NewDistances(points []dimension.Distancer) []Distance {
	prefix := "x"
	n := len(points)
	distances := []Distance{}

	for i := 0; i < n; i++ {
		// create a new fixed cluster
		distance := Distance{
			Set: set.NewSet(fmt.Sprintf("%s%d", prefix, i+1)),
		}

		// if we reached the end of the list
		// this means we are on the last distance
		if i+1 == n {
			distances = append(distances, distance)
			continue
		}

		distance.Points = make(map[set.Set]float64)
		// for every left point compute the distance from the
		// fixed cluster to other clusters
		for j := i + 1; j < n; j++ {
			// length is the distance between cluster i and j
			length := points[i].Distance(points[j])
			// create the cluster name  for j
			key := set.NewSet(fmt.Sprintf("%s%d", prefix, j+1))
			// assign distance to the map
			distance.Points[key] = length
		}
		distances = append(distances, distance)
	}

	return distances
}

// Distance holds all distances from one cluster to all other clusters
type Distance struct {
	// Cluster the fixed point
	Set set.Set
	// Points mapping between the cluster and
	// all other clusters based on the distance
	Points map[set.Set]float64
}

var _ fmt.Stringer = (*Distance)(nil)

// String returns the string representation
// of the cluster distance table
func (d Distance) String() string {
	str := fmt.Sprintf("%s =>", d.Set)
	for cluster, distance := range d.Points {
		str += fmt.Sprintf(" %s:%.2f", cluster, distance)
	}

	str += " "
	return str
}

// Merge merges two sets together and removes
// the distance in the map that has the set given
func (d *Distance) Merge(c set.Set) {
	if d.Set == c {
		return
	}

	d.Set.Add(c)
	delete(d.Points, c)
}

// Best picks a pair of clusters and minimum distance
// of the hole row of distances
// If row does not contain distance points it will return
// an empty pair and 0.0
func (d Distance) Best() (set.Set, set.Set, float64) {
	bestDistance := -1.0
	bestCluster := set.NewSet()
	newDistance := 0.0

	if d.Points == nil || len(d.Points) == 0 {
		return bestCluster, bestCluster, 0
	}

	for cluster, distance := range d.Points {
		if bestDistance == -1.0 {
			bestDistance = distance
			bestCluster = cluster
			continue
		}
		newDistance = math.Min(bestDistance, distance)
		if newDistance != bestDistance {
			bestCluster = cluster
			bestDistance = newDistance
			continue
		}

		if newDistance == bestDistance {
			if cluster.Priority(bestCluster) {
				bestCluster = cluster
			}
		}
	}

	d.Set.Add(bestCluster)
	return d.Set, bestCluster, bestDistance
}
