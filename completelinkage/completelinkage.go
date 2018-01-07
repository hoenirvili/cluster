// Package providing basic semantics for choosing different
// clusters that are best fitted for complete linkage clustering
package completelinkage

import (
	"math"

	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
)

// CompleteLinkage type that represents the single linkage
// bottom up cluster semantics
type CompleteLinkage struct{}

// NewCompleteLinkage creates a new CompleteLinkage pointer
func NewCompleteLinkage() *CompleteLinkage {
	return &CompleteLinkage{}
}

// Swap swaps the first distance with the second distance
// if the first distance is greater than the second one
// If not it does viceversa
func (c CompleteLinkage) Swap(first, second distance.Distance) {
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
func (c CompleteLinkage) Recompute(based set.Set, on map[set.Set]float64) (float64, []set.Set) {
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
