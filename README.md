# cluster

Simple bottom up agglomerative cluster library that supports one and two dimensional points. 
For computing the distance between two dimensional points we use the Euclidian distance.

## Api example


[![Build Status](https://travis-ci.org/hoenirvili/cluster.svg?branch=master)](https://travis-ci.org/hoenirvili/cluster) [![GoDoc](https://godoc.org/github.com/hoenirvili/cluster?status.svg)](https://godoc.org/github.com/hoenirvili/cluster) [![Go Report Card](https://goreportcard.com/badge/github.com/hoenirvili/cluster)](https://goreportcard.com/report/github.com/hoenirvili/cluster) [![Coverage Status](https://coveralls.io/repos/github/hoenirvili/cluster/badge.svg?branch=master)](https://coveralls.io/github/hoenirvili/cluster?branch=master)


#### One dimensional point


```go
package main

import (

    "fmt"

    "github.com/hoenirvili/cluster"
    "github.com/hoenirvili/cluster/dimension/one"
    "github.com/hoenirvili/cluster/distance"

)

func main() {
    points := one.NewDistances(-0.3, 0.1, 0.2, 0.4, 1.6, 1.7, 1.9, 2.0)
    distances := distance.NewDistances(points)
    k := 3
    clusters := cluster.Fit(distances, cluster.SingleLinkage, k)
    fmt.Println(clusters)
}

```


#### Two dimensional point


```go
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
    k := 3
    distances := distance.NewDistances(points)
    clusters := cluster.Fit(distances, cluster.SingleLinkage, k)
    fmt.Println(clusters)
}

```
