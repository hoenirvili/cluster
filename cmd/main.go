package main

import (
	"fmt"

	"github.com/hoenirvili/cluster"
	"github.com/hoenirvili/cluster/dimension/two"
	"github.com/hoenirvili/cluster/distance"
)

func main() {
	points := two.NewDistances(
		[]float64{-4, -3, -2, -1, 1, 1, 2, 3, 3, 4},
		[]float64{-2, -2, -2, -2, -1, 1, 3, 2, 4, 3},
	)
	distances := distance.NewDistances(points)
	for _, d := range distances {
		fmt.Println(d)
	}
	clusters := cluster.Fit(distances, cluster.SingleLinkage, 6)
	fmt.Println(clusters)
}
