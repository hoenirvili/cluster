package completelinkage

import (
	"math"

	"github.com/hoenirvili/cluster/set"
)

type CompleteLinkge struct{}

func (c CompleteLinkge) Swap(first, second map[set.Set]float64) {
	if first == nil || second == nil ||
		len(first) == 0 || len(second) == 0 {
		return
	}

	for fc, f := range first {
		s, ok := second[fc]
		if !ok {
			continue
		}
		first[fc] = math.Max(f, s)
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
