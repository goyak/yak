package cmd

import (
	"fmt"
	"os"
)

func YakRoot() string {
	root := os.Getenv("YAKPATH")
	if root == "" {
		root = fmt.Sprintf("%s/yak", os.Getenv("HOME"))
	}
	return root
}
