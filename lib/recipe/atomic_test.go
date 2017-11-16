package recipe

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// some code here to check arguments perhaps?
	fmt.Fprintf(os.Stdout, dockerRunResult)
	os.Exit(0)
}

const dockerRunResult = `{
    "deployments": [
        {
            "base-checksum": "aa467cfa26d8c9914275c692219ec018c1f31dfc036b1bf4fa1d3ab227a2a8d0",
            "base-commit-meta": {
                "rpmostree.clientlayer": true,
                "rpmostree.clientlayer_version": 2,
                "rpmostree.packages": [
                    "NetworkManager-l2tp-gnome",
                    "fcitx",
                    "fcitx-chewing",
                    "git",
                    "git-lfs",
                    "gnome-tweak-tool",
                    "gparted",
                    "snapd",
                    "virt-manager"
                ],
                "rpmostree.removed-base-packages": [],
                "rpmostree.replaced-base-packages": [],
                "rpmostree.state-sha512": "c0c8bfce82fa2a3b1d5ca654547aeb0244a128cd517b1efe88c3b6b5e0337bc65ad254b9d3bc10676dce566668e3ed286d9fd435b4134c55aa9eeedd81e230f9",
                "version": "Rawhide.20171106.n.0"
            },
            "base-local-replacements": [],
            "base-removals": [],
            "base-timestamp": 1509962363,
            "base-version": "Rawhide.20171106.n.0",
            "booted": true,
            "checksum": "9205b7d928029aa007e4f98f94e1c1ddbe63b8c2e980b9a8fd4aecfa16e1ea20",
            "gpg-enabled": false,
            "id": "fedora-9205b7d928029aa007e4f98f94e1c1ddbe63b8c2e980b9a8fd4aecfa16e1ea20.0",
            "layered-commit-meta": {
                "rpmostree.clientlayer": true,
                "rpmostree.clientlayer_version": 2,
                "rpmostree.packages": [
                    "NetworkManager-l2tp-gnome",
                    "fcitx",
                    "fcitx-chewing",
                    "git",
                    "git-lfs",
                    "gnome-tweak-tool",
                    "gparted",
                    "snapd",
                    "virt-manager"
                ],
                "rpmostree.removed-base-packages": [],
                "rpmostree.replaced-base-packages": [],
                "rpmostree.state-sha512": "c0c8bfce82fa2a3b1d5ca654547aeb0244a128cd517b1efe88c3b6b5e0337bc65ad254b9d3bc10676dce566668e3ed286d9fd435b4134c55aa9eeedd81e230f9",
                "version": "Rawhide.20171106.n.0"
            },
            "origin": "fedora-ws-rawhide:fedora/rawhide/x86_64/workstation",
            "osname": "fedora",
            "packages": [
                "NetworkManager-l2tp-gnome",
                "fcitx",
                "fcitx-chewing",
                "git",
                "git-lfs",
                "gnome-tweak-tool",
                "gparted",
                "snapd",
                "virt-manager"
            ],
            "regenerate-initramfs": false,
            "requested-base-local-replacements": [],
            "requested-base-removals": [],
            "requested-local-packages": [],
            "requested-packages": [
                "NetworkManager-l2tp-gnome",
                "grub2-efi-x64",
                "grub2-efi",
                "fcitx",
                "git-lfs",
                "git",
                "gnome-tweak-tool",
                "virt-manager",
                "fcitx-chewing",
                "gparted",
                "snapd"
            ],
            "serial": 0,
            "timestamp": 1510543599,
            "unlocked": "none",
            "version": "Rawhide.20171106.n.0"
        },
        {
            "base-checksum": "aa466cfa26d8c9914275c692219ec018c1f31dfc036b1bf4fa1d3ab227a2a8d0",
            "base-commit-meta": {
                "rpmostree.clientlayer": true,
                "rpmostree.clientlayer_version": 2,
                "rpmostree.packages": [
                    "NetworkManager-l2tp-gnome",
                    "fcitx",
                    "fcitx-chewing",
                    "git",
                    "git-lfs",
                    "gnome-tweak-tool",
                    "gparted",
                    "snapd",
                    "virt-manager"
                ],
                "rpmostree.removed-base-packages": [],
                "rpmostree.replaced-base-packages": [],
                "rpmostree.state-sha512": "fbc04a91a9d9436861f2f49315daebf6f0bb25a316b885f89eaa05b25c47844fa7fd5efe0df66001e3004bd98612828893fdee1ff7d8c6983011c7ebbd4f8f62",
                "version": "Rawhide.20171106.n.0"
            },
            "base-local-replacements": [],
            "base-removals": [],
            "base-timestamp": 1509962363,
            "base-version": "Rawhide.20171106.n.0",
            "booted": false,
            "checksum": "db9af884cb241c11abfb334b630c57bb0857817eed5dbeda68fc245418557fd8",
            "gpg-enabled": false,
            "id": "fedora-db9af884cb241c11abfb334b630c57bb0857817eed5dbeda68fc245418557fd8.0",
            "layered-commit-meta": {
                "rpmostree.clientlayer": true,
                "rpmostree.clientlayer_version": 2,
                "rpmostree.packages": [
                    "NetworkManager-l2tp-gnome",
                    "fcitx",
                    "fcitx-chewing",
                    "git",
                    "git-lfs",
                    "gnome-tweak-tool",
                    "gparted",
                    "snapd",
                    "virt-manager"
                ],
                "rpmostree.removed-base-packages": [],
                "rpmostree.replaced-base-packages": [],
                "rpmostree.state-sha512": "fbc04a91a9d9436861f2f49315daebf6f0bb25a316b885f89eaa05b25c47844fa7fd5efe0df66001e3004bd98612828893fdee1ff7d8c6983011c7ebbd4f8f62",
                "version": "Rawhide.20171106.n.0"
            },
            "origin": "fedora-ws-rawhide:fedora/rawhide/x86_64/workstation",
            "osname": "fedora",
            "packages": [
                "NetworkManager-l2tp-gnome",
                "fcitx",
                "fcitx-chewing",
                "git",
                "git-lfs",
                "gnome-tweak-tool",
                "gparted",
                "snapd",
                "virt-manager"
            ],
            "regenerate-initramfs": false,
            "requested-base-local-replacements": [],
            "requested-base-removals": [],
            "requested-local-packages": [],
            "requested-packages": [
                "NetworkManager-l2tp-gnome",
                "grub2-efi-x64",
                "grub2-efi",
                "fcitx",
                "git-lfs",
                "git",
                "gnome-tweak-tool",
                "virt-manager",
                "fcitx-chewing",
                "gparted",
                "snapd"
            ],
            "serial": 0,
            "timestamp": 1510284912,
            "unlocked": "none",
            "version": "Rawhide.20171106.n.0"
        }
    ],
    "transaction": null
}`

func TestGetCurrentChecksum(t *testing.T) {
	expected := "aa467c"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	checksum := getCurrentChecksum()

	if checksum != expected {
		t.Errorf("Expected %q, got %q", expected, checksum)
	}
}

func TestAtomicRecipeConfig(t *testing.T) {
	recipe := LoadRecipeConfig("testdata/atomic.yml")
	recipeType := fmt.Sprintf("%T", recipe)
	if recipeType != "recipe.AtomicRecipeConfig" {
		t.Error("The recipeType is not recipe.AtomicRecipeConfig")
	}
}
