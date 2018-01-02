package averagelinkage

import (
	"math"

	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
)

type AverageLinkage struct {
	// this will hold a copy of the table of distances
	// in the 0 interation
	Table []distance.Distance
}

func (a AverageLinkage) Swap(first, second map[set.Set]float64) {
	panic("This is not implemented yet")
	if first == nil || second == nil ||
		len(first) == 0 || len(second) == 0 {
		return
	}

	for fc, f := range first {
		s, ok := second[fc]
		if !ok {
			continue
		}
		//TODO(hoenir): fix this.
		first[fc] = math.Min(f, s)
	}
}

func (a AverageLinkage) Recompute(base set.Set, points map[set.Set]float64) (float64, []set.Set) {
	panic("This is not implemented yet")
	return 0, []set.Set{}
}
