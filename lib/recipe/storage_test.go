package recipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadRecipe(t *testing.T) {
	r := LoadRecipe(".", "testdata/foo")

	assert.IsType(t, new(IRecipeConfig), &r)
}
