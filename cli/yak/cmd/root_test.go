package cmd

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type YakSuite struct {}

var _ = Suite(&YakSuite{})

func mockArgs(args ...string) (restore func()) {
        old := os.Args
        os.Args = args
        return func() { os.Args = old }
}

func (s *YakSuite) TestUnknownCommentErrorResult(c *C) {
	restore := mockArgs("yak", "foo")
	defer restore()

        err := RootCmd.Execute()
        c.Assert(err, ErrorMatches, `unknown command \"foo\" for \"yak\"`)
}
