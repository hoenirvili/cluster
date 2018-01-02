// Clustering is one of several methods of hierarchical clustering.
// At the beginning of the process, each element is in a cluster of its own.
// The clusters are then sequentially combined into larger clusters
// until all elements end up being in the same cluster.

package cluster

import (
	"fmt"

	"github.com/hoenirvili/cluster/averagelinkage"
	"github.com/hoenirvili/cluster/completelinkage"
	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
	"github.com/hoenirvili/cluster/singlelinkage"
)

// strategy represents the type used
// in refitting the table of distances
type strategy uint8

const (
	// SingleLinkage it is based on grouping clusters
	// in bottom-up fashion (agglomerative clustering), at each step
	// combining two clusters that contain the closest pair of elements
	// not yet belonging to the same cluster as each other.
	SingleLinkage strategy = iota
	// CompleteLinkage it is based on grouping clusters
	// in bottom-up fashion (agglomerative clustering), at each step
	// combining two clusters that contain the farthest pair of elements
	// not yet belonging to the same cluster as each other.
	CompleteLinkage
	// AverageLinkage hierarchical clustering, the distance between two clusters
	// is defined as the average distance between each point in one cluster
	// to every point in the other cluster. For example, the distance between clusters
	//  “r” and “s” to the left is equal to the average
	// length each arrow between connecting the points of one cluster to the other.
	AverageLinkage
)

// Swapper defines the criteria of which we swap and
// replace the best point based on the cluster implementation
// and recompute their distances
type Swapper interface {
	Swap(first, second map[set.Set]float64)
	Recompute(based set.Set, on map[set.Set]float64) (float64, []set.Set)
}

func refit(points *[]distance.Distance, first, second set.Set, s Swapper) {
	if first == second {
		return
	}

	n, j := len(*points), 0
	var base set.Set
	for i := 0; i < n; i++ {
		// we found the first cluster
		if (*points)[i].Set == first {
			base = (*points)[i].Set
			for j = 0; j < n; j++ {
				if i == j {
					continue
				}
				// we found the second cluster
				if (*points)[j].Set == second {
					s.Swap((*points)[i].Points, (*points)[j].Points)
					break
				}
			}
			break
		}
	}

	// we didn't find the second cluster
	if j != n {
		*points = append((*points)[:j], (*points)[j+1:]...)
	}

	// the last should always be nil
	last := len(*points) - 1
	if (*points)[last].Points != nil {
		for key := range (*points)[last].Points {
			delete((*points)[last].Points, key)
		}
		(*points)[last].Points = nil
	}

	recomputeDistances(*points, base, s)
}

// recomputeDistances recomputes the table of distances using
// the base cluster as relative distances.
func recomputeDistances(points []distance.Distance, base set.Set, s Swapper) {
	n := len(points)
	for i := 0; i < n; i++ {
		if points[i].Set == base {
			continue
		}

		best, toDelete := s.Recompute(base, points[i].Points)
		if best == -1.0 {
			continue
		}

		for _, key := range toDelete {
			delete(points[i].Points, key)
		}

		for cluster := range points[i].Points {
			if base.In(cluster) {
				value := points[i].Points[cluster]
				delete(points[i].Points, cluster)
				key := base
				points[i].Points[key] = value
				break
			}
		}
	}
}

// Fit will cluster the points into k clusters based on the strategy provided
func Fit(points []distance.Distance, s strategy, k int) []set.Set {
	if k <= 0 || k > len(points) {
		return nil
	}

	var swapper Swapper
	switch s {
	case SingleLinkage:
		swapper = &singlelinkage.SingleLinkage{}
	case CompleteLinkage:
		swapper = &completelinkage.CompleteLinkge{}
	case AverageLinkage:
		swapper = &averagelinkage.AverageLinkage{Table: points}
	}

	pair := struct{ first, second set.Set }{}
	for n := len(points); k != n; n = len(points) {
		bestDistance := -1.0
		j := 0
		for i := 0; i < n; i++ {
			f, s, distance := points[i].Best()
			if f == s {
				continue
			}
			if bestDistance == -1 || bestDistance > distance {
				bestDistance = distance
				pair.first, pair.second = f, s
				j = i
			}
		}

		points[j].Merge(pair.second)
		refit(&points, pair.first, pair.second, swapper)

		fmt.Println("BEGIN")
		for _, d := range points {
			fmt.Println(d)
		}
		fmt.Println("END")
		fmt.Println()
	}

	cls := make([]set.Set, 0, k)
	for _, c := range points {
		cls = append(cls, c.Set)
	}

	return cls
}
