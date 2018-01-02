package main

import (
	"fmt"

	"github.com/hoenirvili/cluster"
	"github.com/hoenirvili/cluster/dimension/one"
)

func main() {
	points := one.NewDistances(-0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0)
	distances := cluster.NewDistances(points)
	for _, d := range distances {
		fmt.Println(d)

	}
	clusters := cluster.Fit(distances, cluster.SingleLinkage, 4)
	fmt.Println(clusters)
}
