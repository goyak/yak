package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Hook up gocheck into the "go test" runner.

func mockArgs(args ...string) (restore func()) {
	old := os.Args
	os.Args = args
	return func() { os.Args = old }
}

func TestUnknownCommentErrorResult(t *testing.T) {
	restore := mockArgs("yak", "foo")
	defer restore()

	err := RootCmd.Execute()
	// assert.Equal(t, ErrorMatches, `unknown command \"foo\" for \"yak\"`)
	assert.Regexp(t, `unknown command \"foo\" for \"yak\"`, err.Error())

	// c.Assert(err, ErrorMatches, `unknown command \"foo\" for \"yak\"`)
}
