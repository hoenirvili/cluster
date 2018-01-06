package completelinkage

import (
	"math"

	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
)

type CompleteLinkge struct{}

// NewCompleteLinakge creates a new CompleteLinkage pointer
func NewCompleteLinkage() *CompleteLinkge {
	return &CompleteLinkge{}
}

// Swap swaps the first distance with the second distance
// if the first distance is greater than the second one
// If not it does viceversa
func (c CompleteLinkge) Swap(first, second distance.Distance) {
	for fc, f := range first.Points {
		s, ok := second.Points[fc]
		if !ok {
			continue
		}
		first.Points[fc] = math.Max(f, s)
	}
}

// Recompute recomputes the remaining distances after
// the swap process is done based on the cluster provided and returns the best
// distance alongside with the keys of the map of distances that should be removed
func (s CompleteLinkge) Recompute(based set.Set, on map[set.Set]float64) (float64, []set.Set) {
	toBeDeleted := []set.Set{}
	previous := set.NewSet()
	best := -1.0
	for cluster, distance := range on {
		if based.In(cluster) {
			if best == -1.0 {
				best = distance
				previous = cluster
				continue
			}
			if best > distance {
				toBeDeleted = append(toBeDeleted, cluster)
			} else {
				best = distance
				toBeDeleted = append(toBeDeleted, previous)
				previous = cluster
			}
		}
	}

	return best, toBeDeleted
}
