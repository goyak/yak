package recipe

import (
	"fmt"
	"os/exec"
	"strings"
)

type AtomicRecipeConfig struct {
	BaseRecipeConfig
}

func (r AtomicRecipeConfig) IsInstallable() bool {
	return false
}

func (r AtomicRecipeConfig) Install() bool {
	remoteName := strings.Split(r.Branch, "/")[0]
	fmt.Printf(">>> install : n:%s s:%s c:%s \n", r.Name, r.Source, r.Commit)
	// ostree remote add --if-not-exists --no-gpg-verify fedora-atomic-26 https://kojipkgs.fedoraproject.org/atomic/26
	addRemoteCmd := exec.Command("ostree", "remote", "add", "--if-not-exists", "--no-gpg-verify", remoteName, r.Source)
	fmt.Printf("::: %s  %q \n", addRemoteCmd.Path, addRemoteCmd.Args)
	// addRemoteCmd.Run()

	pullCmd := exec.Command("ostree", "pull", remoteName, r.Commit)
	fmt.Printf("::: %s  %q \n", pullCmd.Path, pullCmd.Args)
	// pullCmd.Run()

	// ostree admin deploy d518b37c348eb814093249f035ae852e7723840521b4bcb4a271a80b5988c44a
	// rpm-ostree deploy 173278f2ccba80c5cdda4b9530e6f0388177fb6d27083

	cmd := exec.Command("ostree", "admin", "deploy", r.Commit)
	cmd.Run()
	fmt.Printf("::: %s  %q \n", cmd.Path, cmd.Args)

	// Prepare to reboot
	// cmd := exec.Command("ostree", "/tmp/aaa")
	// cmd.Run()
	return true
}
