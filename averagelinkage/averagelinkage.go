package averagelinkage

import (
	"github.com/hoenirvili/cluster/distance"
	"github.com/hoenirvili/cluster/set"
)

// AverageLinkage type that represents the single linkage
// bottom up cluster semantics
type AverageLinkage struct {
	// this will hold a copy of the table of distances
	// in the 0 interation
	Table []distance.Distance
}

// NewAverageLinkage copies the table of distances provided and returns
// a new pointer to AverageLinkage
// If the table is nil, or empty it will return nil
func NewAverageLinkage(table []distance.Distance) *AverageLinkage {
	if table == nil || len(table) == 0 {
		return nil
	}

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

// setEq tests if two sets/clusters are equal
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

// distance computes the distance between two clusters from the
// original table that has been copied
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

// distances returns a list of distances based on the points
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

// average returns the average of all distances
func (a AverageLinkage) average(distances []float64) float64 {
	n := len(distances)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += distances[i]
	}

	return sum / float64(n)
}

// Swap swaps the first distance with the second distance
// using the average distance of a cluster point
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

// Recompute recomputes the remaining distances after
// the swap process is done based on the cluster provided and returns the best
// distance alongside with the keys of the map of distances that should be removed
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
