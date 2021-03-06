package set_test

import (
	"github.com/hoenirvili/cluster/set"
	gc "gopkg.in/check.v1"
)

type setSuite struct{}

var _ = gc.Suite(&setSuite{})

func (cs setSuite) newSet(c *gc.C) set.Set {
	cls := set.NewSet("x2", "x3", "x1")
	c.Assert(cls, gc.NotNil)
	c.Assert(cls, gc.DeepEquals, set.Set("x1,x2,x3"))
	return cls
}

func (cs setSuite) TestSlice(c *gc.C) {
	cluster := cs.newSet(c)
	slice := cluster.Slice()
	c.Assert(slice, gc.DeepEquals, []string{"x1", "x2", "x3"})
}

func (cs setSuite) TestEmpty(c *gc.C) {
	cluster := cs.newSet(c)
	empty := cluster.Empty()
	c.Assert(empty, gc.Equals, false)

	cluster = set.Set("")
	empty = cluster.Empty()
	c.Assert(empty, gc.Equals, true)
}

func (cs setSuite) TestIn(c *gc.C) {
	cls := cs.newSet(c)
	found := cls.In("x1")
	c.Assert(found, gc.Equals, true)

	found = cls.In("x1,x2,x3")
	c.Assert(found, gc.Equals, true)

	found = cls.In("x2,x3")
	c.Assert(found, gc.Equals, true)

	found = cls.In("x7")
	c.Assert(found, gc.Equals, false)

	found = cls.In("x8,x1,x2,x9")
	c.Assert(found, gc.Equals, false)
}

func (cs setSuite) TestSimple(c *gc.C) {
	cls := cs.newSet(c)
	c.Assert(cls.Simple(), gc.Equals, false)

	cls = set.Set("x1")
	c.Assert(cls.Simple(), gc.Equals, true)
}

func (cs setSuite) TestPriority(c *gc.C) {
	one := set.Set("x2")
	second := set.Set("x1")

	p := one.Priority(second)
	c.Assert(p, gc.Equals, false)

	one, second = second, one

	p = one.Priority(second)
	c.Assert(p, gc.Equals, true)
}

func (cs setSuite) TestLen(c *gc.C) {
	cls := cs.newSet(c)
	got := cls.Len()
	expected := 3
	c.Assert(got, gc.Equals, expected)

	cls = set.NewSet()
	got = cls.Len()
	expected = 0
	c.Assert(got, gc.Equals, expected)
}

func (cs setSuite) TestSwap(c *gc.C) {
	cls := cs.newSet(c)
	cls.Swap(0, 2)
	slice := cls.Slice()
	c.Assert(slice, gc.DeepEquals, []string{"x3", "x2", "x1"})
}

func (cs setSuite) TestLess(c *gc.C) {
	cls := cs.newSet(c)
	got := cls.Less(0, 2)
	c.Assert(got, gc.Equals, true)
}

func (cs setSuite) TestAdd(c *gc.C) {
	cls := cs.newSet(c)
	point := set.Set("x5")
	cls.Add(point)
	expected := set.Set("x1,x2,x3," + point)
	c.Assert(cls, gc.DeepEquals, expected)
}

func (cs setSuite) TestAddFromEmpty(c *gc.C) {
	cls := set.NewSet()
	point := set.Set("x5")
	cls.Add(point)
	c.Assert(cls, gc.DeepEquals, point)
	slice := cls.Slice()
	c.Assert(slice, gc.DeepEquals, []string{"x5"})
	c.Assert(len(slice), gc.Equals, cls.Len())

	cls = set.NewSet()
	point1, point2 := set.Set("x5"), set.Set("x6")
	cls.Add(point1)
	cls.Add(point2)
	c.Assert(cls, gc.DeepEquals, set.Set("x5,x6"))
	c.Assert(cls.Len(), gc.Equals, len(cls.Slice()))
}

func (cs setSuite) TestAddDuplicate(c *gc.C) {
	cls := cs.newSet(c)
	point := set.Set("x5")
	cls.Add(point)
	expected := set.Set("x1,x2,x3," + point)
	c.Assert(cls, gc.DeepEquals, expected)
	cls.Add(point)
	c.Assert(cls, gc.DeepEquals, expected)
}

func (cs setSuite) TestDelete(c *gc.C) {
	cls := cs.newSet(c)
	point := set.Set("x3")
	cls.Delete(point)
	expected := set.Set("x1,x2")
	c.Assert(cls, gc.DeepEquals, expected)
}

func (cs setSuite) TestDeleteWhenEmpty(c *gc.C) {
	cls := cs.newSet(c)
	for _, item := range []set.Set{"x1", "x2", "x3"} {
		cls.Delete(item)
	}

	empty := cls.Empty()
	c.Assert(empty, gc.Equals, true)

	cls.Delete("x1")

	empty = cls.Empty()
	c.Assert(empty, gc.Equals, true)
}

func (cs setSuite) TestString(c *gc.C) {
	expected := "{x1,x2,x3}"
	cluster := cs.newSet(c)
	got := cluster.String()
	c.Assert(got, gc.DeepEquals, expected)
}
