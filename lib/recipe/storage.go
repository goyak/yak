package recipe

import (
	"os"

	"github.com/goyak/yak/lib/env"
)

// dir.go  docker.go  ostree.go  repo.go  shortcut.go  tarball.go
func isDir(pth string) (bool, error) {
	fi, err := os.Stat(pth)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func LoadRecipe(yakroot string, repo string) IRecipeConfig {
	path := yakroot + "/" + env.LocalDbDir + "/" + repo
	result, _ := isDir(path)
	if result {
		return LoadRecipeConfig(path + "/" + "yak.yml")
	}
	r := BaseRecipeConfig{}
	r.Repo = repo
	return r
}
