package recipe

import (
	"fmt"
	"testing"
)

func TestAtomicRecipe(t *testing.T) {
	recipe := LoadRecipe("atomic_yaml_testdata")
	recipeType := fmt.Sprintf("%T", recipe)
	if recipeType != "recipe.AtomicRecipe" {
		t.Error("The recipeType is not recipe.AtomicRecipe")
	}
}
