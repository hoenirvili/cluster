package averagelinkage

import "github.com/hoenirvili/cluster/set"

type AverageLinkage struct{}

func (a AverageLinkage) Swap(first, second map[set.Set]float64) {
	panic("This is not implemented yet")
}

func (a AverageLinkage) Recompute(base set.Set, points map[set.Set]float64) (float64, []set.Set) {
	panic("This is not implemented yet")
	return 0, []set.Set{}
}
