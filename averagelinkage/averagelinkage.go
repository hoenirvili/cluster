package averagelinkage

import (
	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
)

type AverageLinkage struct {
	// this will hold a copy of the table of distances
	// in the 0 interation
	Table []distance.Distance
}

func NewAverageLinkage(table []distance.Distance) *AverageLinkage {
	a := &AverageLinkage{}
	n := len(table)
	a.Table = make([]distance.Distance, n, n)

	for key, row := range table {
		a.Table[key].Set = row.Set
		a.Table[key].Points = make(map[set.Set]float64, len(row.Points))
		for mkey, col := range row.Points {
			a.Table[key].Points[mkey] = col
		}
	}

	return a
}

func (a AverageLinkage) setEq(row, col set.Set) bool {
	if row == col {
		return true
	}

	if row.In(col) {
		return true
	}

	if col.In(row) {
		return true
	}

	return false
}

func (a AverageLinkage) distance(row, col set.Set) float64 {
	for _, rowTable := range a.Table {
		if a.setEq(rowTable.Set, row) {
			for colTable, val := range rowTable.Points {
				if a.setEq(colTable, col) {
					return val
				}
			}
		}
	}

	panic("row, col pair given distance not found")
}

func (a AverageLinkage) distances(points []string) []float64 {
	n := len(points)
	distances := []float64{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			row, col := set.Set(points[i]), set.Set(points[j])
			distances = append(distances, a.distance(row, col))
		}
	}
	return distances
}

func (a AverageLinkage) average(distances []float64) float64 {
	n := len(distances)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += distances[i]
	}

	return sum / float64(n)
}

func (a AverageLinkage) Swap(first, second distance.Distance) {
	points := first.Set.Slice()
	n := len(points)

	for fc := range first.Points {
		if _, ok := second.Points[fc]; !ok {
			continue
		}
		if n == len(points) {
			points = append(points, string(fc))
		} else {
			points[n] = string(fc)
		}

		distances := a.distances(points)
		first.Points[fc] = a.average(distances)
	}
}

func (a AverageLinkage) Recompute(based set.Set, on map[set.Set]float64) (float64, []set.Set) {
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
