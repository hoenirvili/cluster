package completelinkage

import (
	"math"

	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
)

type CompleteLinkge struct{}

func NewCompleteLinkage() *CompleteLinkge {
	return &CompleteLinkge{}
}

func (c CompleteLinkge) Swap(first, second distance.Distance) {
	for fc, f := range first.Points {
		s, ok := second.Points[fc]
		if !ok {
			continue
		}
		first.Points[fc] = math.Max(f, s)
	}
}

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
