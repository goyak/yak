package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYakRoot(t *testing.T) {
	yakroot := "/tmp/yak"
	os.Setenv("YAKROOT", yakroot)

	assert.Equal(t, YakRoot(), yakroot)
}

func TestYakRootDefault(t *testing.T) {
	os.Unsetenv("YAKROOT")
	os.Setenv("HOME", "/p")

	assert.Equal(t, YakRoot(), "/p/yak")
}
