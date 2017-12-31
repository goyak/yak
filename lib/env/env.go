package env

import (
	"fmt"
	"os"
)

const LocalIndex = `db/local.yml`
const LocalDbDir = `db`
const DataDir = `data`

func YakRoot() string {
	root := os.Getenv("YAKPATH")
	if root == "" {
		root = fmt.Sprintf("%s/yak", os.Getenv("HOME"))
	}
	return root
}
