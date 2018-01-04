package singlelinkage

import (
	"math"

	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
)

type SingleLinkage struct{}

func NewSingleLinkage() *SingleLinkage {
	return &SingleLinkage{}
}

func (s SingleLinkage) Swap(first, second distance.Distance) {
	for fc, f := range first.Points {
		s, ok := second.Points[fc]
		if !ok {
			continue
		}
		first.Points[fc] = math.Min(f, s)
	}
}

func (s SingleLinkage) Recompute(based set.Set, on map[set.Set]float64) (float64, []set.Set) {
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
