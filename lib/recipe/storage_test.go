package recipe

import (
	"testing"
)

func TestLoadRecipe(t *testing.T) {
	r := LoadRecipe("testdata/foo")
	r.Dump()
}
