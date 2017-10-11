package recipe

import (
	"os"
)

// dir.go  docker.go  ostree.go  repo.go  shortcut.go  tarball.go
func isDir(pth string) (bool, error) {
	fi, err := os.Stat(pth)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func LoadRecipe(str string) IRecipeConfig {
	// if str == dir
	result, _ := isDir(str)
	if result {
		return LoadRecipeConfig(str + "/" + "yak.yml")
	}
	return nil
}
