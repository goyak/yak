package recipe

import (
	"os/exec"
)

type AtomicRecipeConfig struct {
	BaseRecipeConfig
}

func (r AtomicRecipeConfig) IsInstallable() bool {
	return false
}

func (r AtomicRecipeConfig) Install() bool {
	cmd := exec.Command("touch", "/tmp/aaa")
	cmd.Run()
	return true
}
