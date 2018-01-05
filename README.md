# cluster



## Api example


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
    distances := distance.NewDistances(points)
    clusters := cluster.Fit(distances, cluster.SingleLinkage, k)
    fmt.Println(clusters)
}

```