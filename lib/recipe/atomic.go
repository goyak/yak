package recipe

type AtomicRecipe struct {
	Recipe
}

func (r AtomicRecipe) IsInstallable() bool {
	return false
}
