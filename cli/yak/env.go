package main

import (
	"fmt"
	"os"
)

func YakRoot() string {
	root := os.Getenv("YAKROOT")
	if root == "" {
		root = fmt.Sprintf("%s/yak", os.Getenv("HOME"))
	}
	return root
}
