package util_test

import (
	"github.com/hoenirvili/cluster/util"
	gc "gopkg.in/check.v1"
)

type utilSuite struct{}

var _ = gc.Suite(&utilSuite{})

func (u utilSuite) TestRound(c *gc.C) {
	f := 7.2362
	f = util.Round(f, 2)
	c.Assert(f, gc.Equals, 7.24)

	f = util.Round(0, 10)
	c.Assert(f, gc.Equals, 0.0)
}
