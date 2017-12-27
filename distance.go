package cluster

import (
	"fmt"

	"github.com/hoenirvili/cluster/dimension"
)

// Distance holds all distances from one cluster to all other clusters
type Distance struct {
	// Cluster the fixed point
	Cluster Cluster
	// Points mapping between the cluster and
	// all other clusters based on the distance
	Points map[Cluster]float64
}

// String returns the string representation
// of the cluster distance table
func (d Distance) String() string {
	str := fmt.Sprintf("%s =>", d.Cluster)
	for cluster, distance := range d.Points {
		str += fmt.Sprintf(" %s:%.2f", cluster, distance)
	}

	str += " "
	return str
}

// trefit represents the type used
// in refitting the table of distances
type trefit uint8

const (
	SingleLinkage trefit = iota
	CompleteLinkage
	AverageLinkage
)

func singleLinkage(first, second map[Cluster]float64) {
	for key, fv := range first {
		sv, ok := second[key]
		if !ok {
			continue
		}
		if fv > sv {
			first[key] = sv
		} else {
			first[key] = fv
		}
	}
}

// Refit based on the type of refit passed in trefit
// this will refit the hole table based on the best
// provided first and second clusters
func Refit(points *[]Distance, first, second Cluster, t trefit) {
	var compareAndSwap func(first, second map[Cluster]float64)
	switch t {
	case SingleLinkage:
		compareAndSwap = singleLinkage
	case CompleteLinkage:
		fallthrough
	case AverageLinkage:
		panic("This is not implemented yet")
	}

	n := len(*points)
	j := 0 // we save the second cluster to remove it from the table
	for i := 0; i < n; i++ {
		// we found that first cluster
		if (*points)[i].Cluster == first {
			for j = 0; j < n; j++ {
				if i == j {
					continue
				}
				// we found the second cluster
				if (*points)[j].Cluster == second {
					compareAndSwap((*points)[i].Points, (*points)[j].Points)
					break
				}
			}
			break
		}
	}

	*points = append((*points)[:j], (*points)[j+1:]...)
}

// NewDistances returns a table of cluster distances
func NewDistances(points []dimension.Distancer) []Distance {
	prefix := "x"
	n := len(points)
	distances := make([]Distance, 0, 0)

	for i := 0; i < n; i++ {
		// create a new fixed cluster
		distance := Distance{
			Cluster: NewCluster(fmt.Sprintf("%s%d", prefix, i+1)),
		}

		// if we reached the end of the list
		// this means we are on the last distance
		// and we should skip it
		if i+1 == n {
			continue
		}

		distance.Points = make(map[Cluster]float64)
		// for every left point compute the distance from the
		// fixed cluster to other clusters
		for j := i + 1; j < n; j++ {
			// length is the distance between cluster i and j
			length := points[i].Distance(points[j])
			// create the cluster name  for j
			key := NewCluster(fmt.Sprintf("%s%d", prefix, j+1))
			// assign distance to the map
			distance.Points[key] = length
		}
		distances = append(distances, distance)
	}

	return distances
}

// Best returns the best fitted cluster pairs
func (d Distance) Best() (Cluster, Cluster, float64) {
	bestDistance := -1.0
	bestCluster := Cluster("")
	for c, distance := range d.Points {
		if bestDistance == -1.0 {
			bestDistance = distance
			bestCluster = c
			continue
		}

		if bestDistance > distance {
			bestDistance = distance
			bestCluster = c
		}
	}

	return d.Cluster, bestCluster, bestDistance
}

func (d *Distance) Merge(c Cluster) {
	d.Cluster.Add(c)
	delete(d.Points, c)
}
