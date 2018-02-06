// Package singlelinkage providing basic semantics for choosing different
// clusters that are best fitted for single linkage clustering
package singlelinkage

import (
	"math"

	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
)

// SingleLinkage type that represents the single linkage
// bottom up cluster semantics
type SingleLinkage struct{}

// NewSingleLinkage creates a new pointer of type SingleLinkage
func NewSingleLinkage() *SingleLinkage {
	return &SingleLinkage{}
}

// Swap swaps the first distance with the second distance
// if the first distance is lesser than the second one
// If not it does viceversa
func (s SingleLinkage) Swap(first, second distance.Distance) {
	for fc, f := range first.Points {
		s, ok := second.Points[fc]
		if !ok {
			continue
		}
		first.Points[fc] = math.Min(f, s)
	}
}

// Recompute recomputes the remaining distances after
// the swap process is done based on the cluster provided and returns the best
// distance alongside with the keys of the map of distances that should be removed
func (s SingleLinkage) Recompute(based set.Set, on distance.Distance) (float64, []set.Set) {
	toBeDeleted := []set.Set{}
	previous := set.NewSet()
	best := -1.0
	for cluster, distance := range on.Points {
		if based.In(cluster) {
			if best == -1.0 {
				best = distance
				previous = cluster
				continue
			}
			if best > distance {
				best = distance
				toBeDeleted = append(toBeDeleted, previous)
				previous = cluster
			} else {
				toBeDeleted = append(toBeDeleted, cluster)
			}
		}
	}

	return best, toBeDeleted
}
