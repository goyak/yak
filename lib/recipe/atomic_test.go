package recipe

import (
	"fmt"
	"testing"
)

func TestAtomicRecipeConfig(t *testing.T) {
	recipe := LoadRecipeConfig("atomic_yaml_testdata")
	recipeType := fmt.Sprintf("%T", recipe)
	if recipeType != "recipe.AtomicRecipeConfig" {
		t.Error("The recipeType is not recipe.AtomicRecipeConfig")
	}
}
