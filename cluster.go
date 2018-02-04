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
// in fiting the table of distances
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

// swapper defines the criteria of which we swap and
// replace the best point based on the cluster strategy
// implementation and recompute their distances
type swapper interface {
	// Swap swaps the first distance with the second distance
	// based on the cluster algorithm
	Swap(first, second distance.Distance)

	// Recompute recomputes the remaining distances after
	// the swap process is done based on the cluster provided and returns the best
	// distance alongside with the keys of the map of distances that should be removed
	Recompute(based set.Set, on distance.Distance) (best float64, deleted []set.Set)
}

// Fit will fit the points in k clusters based on the strategy of clustering
// provided. This will return the k clusters that best fits the distance points
func Fit(points []distance.Distance, s strategy, k int) []set.Set {
	if k <= 0 || k > len(points) {
		return nil
	}

	// don't modify the original slice, make a copy out of it first
	table := make([]distance.Distance, len(points), len(points))
	for key, row := range points {
		table[key].Set = row.Set
		table[key].Points = make(map[set.Set]float64, len(row.Points))
		for mkey, col := range row.Points {
			table[key].Points[mkey] = col
		}
	}

	var swapper swapper
	switch s {
	case SingleLinkage:
		swapper = singlelinkage.NewSingleLinkage()
	case CompleteLinkage:
		swapper = completelinkage.NewCompleteLinkage()
	case AverageLinkage:
		swapper = averagelinkage.NewAverageLinkage(table)
	}

	pair := struct{ first, second set.Set }{}
	for n := len(table); k != n; n = len(table) {
		bestDistance := -1.0
		j := 0
		for i := 0; i < n; i++ {
			f, s, distance := table[i].Best()
			if f == s {
				continue
			}

			if bestDistance == -1 || bestDistance > distance {
				bestDistance = distance
				pair.first, pair.second = f, s
				j = i
				continue
			}
		}

		table[j].Merge(pair.second)
		table = refit(table, pair.first, pair.second, swapper)

		fmt.Println("BEGIN")
		for _, d := range table {
			fmt.Println(d)
		}
		fmt.Println("END")
	}

	cls := make([]set.Set, 0, k)
	for _, c := range table {
		cls = append(cls, c.Set)
	}

	return cls
}

// refit refits all distance points based on the first and second clusters that has been
// chosen in the i-th iteration. If the clusters provided are the same this will return the same
// points
func refit(points []distance.Distance, first, second set.Set, s swapper) []distance.Distance {
	if first == second {
		return points
	}

	n, j := len(points), 0
	var base set.Set
	for i := 0; i < n; i++ {
		// we found the first cluster
		if points[i].Set == first {
			base = points[i].Set
			for j = 0; j < n; j++ {
				if i == j {
					continue
				}
				// we found the second cluster
				if points[j].Set == second {
					s.Swap(points[i], points[j])
					break
				}
			}
			break
		}
	}

	// we didn't find the second cluster
	if j != n {
		points = append(points[:j], points[j+1:]...)
	}

	// the last should always be nil
	last := len(points) - 1
	if (points)[last].Points != nil {
		for key := range points[last].Points {
			delete(points[last].Points, key)
		}
		points[last].Points = nil
	}

	recomputeDistances(points, base, s)
	return points
}

// recomputeDistances recomputes the table of distances using
// the base cluster as relative distances.
func recomputeDistances(points []distance.Distance, base set.Set, s swapper) {
	n := len(points)
	for i := 0; i < n; i++ {
		if points[i].Set == base {
			continue
		}

		best, toDelete := s.Recompute(base, points[i])
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
