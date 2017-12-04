package env

import (
	"fmt"
	"os"
)

const LocalIndex = `index/local.yml`
const RecipeDir = `recipes`
const IndexDir = `index`

func YakRoot() string {
	root := os.Getenv("YAKPATH")
	if root == "" {
		root = fmt.Sprintf("%s/yak", os.Getenv("HOME"))
	}
	return root
}
