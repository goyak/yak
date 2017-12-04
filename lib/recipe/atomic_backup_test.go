package recipe

import (
	"testing"
)

func TestUpdateBackup(t *testing.T) {
	b := backupDeployment{
		Checksum: "abcd1234",
	}
	b.updateBackup(`/tmp/backup.yaml`)
}
