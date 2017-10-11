package recipe

import (
	"testing"
)

func TestLoadRecipe(t *testing.T) {
	r := LoadRecipe("dir_testdata")
	r.Dump()
}
