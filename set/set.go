// Package set provides a sorted set to represent clusters
// from the given points
package set

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Set describes a set of points with the naming convention
// of x1,x2,x3,x4,x5, x6x7x8, x9x10
// These will be always sorted based on their suffix order
type Set string

var (
	_ fmt.Stringer   = (*Set)(nil)
	_ sort.Interface = (*Set)(nil)
)

// NewSet create a cluster from the given points
func NewSet(points ...string) Set {
	switch len(points) {
	case 0:
		return Set("")
	case 1:
		return Set(points[0])
	}

	set := Set(strings.Join(points, ","))
	sort.Sort(&set)
	return set
}

// Slice returns a slice of string sets
func (s Set) Slice() []string {
	return strings.Split(string(s), ",")
}

// Empty returns true if the set is empty
func (s Set) Empty() bool {
	return s == ""
}

// In returns true if the cluster is found in the set
func (s Set) In(set Set) bool {
	cslice := s.Slice()
	slice := set.Slice()
	n := len(cslice)
	m := len(slice)

	if n == m {
		return s == set
	}

	apparitions := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if slice[i] == cslice[j] {
				apparitions++
			}
		}
	}

	if apparitions == m {
		return true
	}

	return false
}

func (s Set) num(i int) int {
	slice := s.Slice()
	suffix := slice[i][1:]
	num, err := strconv.ParseInt(suffix, 10, 64)
	if err != nil {
		sep := strings.Split(suffix, "x")
		num, err := strconv.ParseInt(sep[0], 10, 64)
		if err != nil {
			panic(err)
		}
		return int(num)
	}

	return int(num)
}

// Simple tests if the set has one element
// this will return true
func (s Set) Simple() bool {
	return 1 == len(s.Slice())
}

// Priority returns true if the given
// cluster is greater than the fixed one
func (s Set) Priority(set Set) bool {
	if !s.Simple() || !set.Simple() {
		return false
	}

	if s.num(0) < set.num(0) {
		return true
	}

	return false
}

// Len returns the number of points in a cluster
func (s Set) Len() int {
	if s == Set("") {
		return 0
	}
	return len(s.Slice())
}

// Swap swaps cluster points
func (s *Set) Swap(i, j int) {
	slice := s.Slice()
	slice[i], slice[j] = slice[j], slice[i]
	*s = Set(strings.Join(slice, ","))
}

// Less compares two cluster points
func (s Set) Less(i, j int) bool { return s.num(i) < s.num(j) }

// Add appends a point in the set
// If the point is already in the set it will add it
func (s *Set) Add(point Set) {
	if s.Len() == 0 {
		*s = point
		return
	}

	ps := string(point)
	if ps == "" {
		return
	}

	slice := s.Slice()
	n := len(slice)

	for i := 0; i < n; i++ {
		if slice[i] == ps {
			return
		}
	}

	slice = append(slice, ps)
	*s = Set(strings.Join(slice, ","))
	sort.Sort(s)
}

// Delete a given point from the set
func (s *Set) Delete(point Set) {
	slice := s.Slice()
	n := len(slice)
	if n == 0 {
		return
	}

	i := -1
	for j := 0; j < n; j++ {
		if slice[j] == string(point) {
			i = j
			break
		}
	}

	if i == -1 {
		return
	}

	slice = append(slice[:i], slice[i+1:]...)
	*s = Set(strings.Join(slice, ","))
	sort.Sort(s)
}

// String returns the cluster representation as a string
func (s Set) String() string {
	return fmt.Sprintf("{%s}", string(s))
}
